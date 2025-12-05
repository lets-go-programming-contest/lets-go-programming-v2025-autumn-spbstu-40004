package conveyer

import (
	"context"
	"fmt"
)

type conveyer interface {
	RegisterDecorator(
		fn func(ctx context.Context, input chan string, output chan string) error,
		input string,
		output string,
	)
	RegisterMultiplexer(
		fn func(ctx context.Context, inputs []chan string, output chan string) error,
		inputs []string,
		output string,
	)
	RegisterSeparator(
		fn func(ctx context.Context, input chan string, outputs []chan string) error,
		input string,
		outputs []string,
	)
	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

type ConveyerImpl struct {
	channels     map[string]chan string
	decorators   []decoratorSpec
	multiplexers []multiplexerSpec
	separators   []separatorSpec
}

type decoratorSpec struct {
	fn     func(ctx context.Context, input chan string, output chan string) error
	input  string
	output string
}

type separatorSpec struct {
	fn      func(ctx context.Context, input chan string, outputs []chan string) error
	input   string
	outputs []string
}

type multiplexerSpec struct {
	fn     func(ctx context.Context, inputs []chan string, output chan string) error
	inputs []string
	output string
}

func New(size int) conveyer {
	return &ConveyerImpl{
		channels: make(map[string]chan string),
	}
}

func (conv *ConveyerImpl) createChannel(id string) {
	if _, exists := conv.channels[id]; !exists {
		conv.channels[id] = make(chan string, 10)
	}
}

func (conv *ConveyerImpl) getChannel(id string) (chan string, error) {
	ch, exists := conv.channels[id]
	if !exists {
		return nil, fmt.Errorf("chan not found")
	}
	return ch, nil
}

func (conv *ConveyerImpl) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	conv.decorators = append(conv.decorators, decoratorSpec{
		fn:     fn,
		input:  input,
		output: output,
	})
}

func (conv *ConveyerImpl) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	conv.multiplexers = append(conv.multiplexers, multiplexerSpec{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
}

func (conv *ConveyerImpl) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	conv.separators = append(conv.separators, separatorSpec{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
}

func (conv *ConveyerImpl) Run(ctx context.Context) error {
	for _, d := range conv.decorators {
		conv.createChannel(d.input)
		conv.createChannel(d.output)
	}

	for _, m := range conv.multiplexers {
		for _, in := range m.inputs {
			conv.createChannel(in)
		}
		conv.createChannel(m.output)
	}

	for _, s := range conv.separators {
		conv.createChannel(s.input)
		for _, out := range s.outputs {
			conv.createChannel(out)
		}
	}

	for _, d := range conv.decorators {
		go func(spec decoratorSpec) {
			inCh, _ := conv.getChannel(spec.input)
			outCh, _ := conv.getChannel(spec.output)
			spec.fn(ctx, inCh, outCh)
		}(d)
	}

	for _, m := range conv.multiplexers {
		go func(spec multiplexerSpec) {
			inputs := make([]chan string, len(spec.inputs))
			for i, id := range spec.inputs {
				ch, _ := conv.getChannel(id)
				inputs[i] = ch
			}
			outCh, _ := conv.getChannel(spec.output)
			spec.fn(ctx, inputs, outCh)
		}(m)
	}

	for _, s := range conv.separators {
		go func(spec separatorSpec) {
			inCh, _ := conv.getChannel(spec.input)
			outputs := make([]chan string, len(spec.outputs))
			for i, id := range spec.outputs {
				ch, _ := conv.getChannel(id)
				outputs[i] = ch
			}
			spec.fn(ctx, inCh, outputs)
		}(s)
	}

	<-ctx.Done()

	for _, ch := range conv.channels {
		close(ch)
	}

	return ctx.Err()
}

func (c *ConveyerImpl) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}
	ch <- data
	return nil
}

func (c *ConveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}
	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return data, nil
}
