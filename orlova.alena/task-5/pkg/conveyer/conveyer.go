package conveyer

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrChan = errors.New("chan not found")
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
	size         int
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

func New(size int) *ConveyerImpl {
	return &ConveyerImpl{
		size:         size,
		channels:     make(map[string]chan string),
		decorators:   make([]decoratorSpec, 0),
		multiplexers: make([]multiplexerSpec, 0),
		separators:   make([]separatorSpec, 0),
	}
}

func (conv *ConveyerImpl) createChannel(id string) {
	if _, exists := conv.channels[id]; !exists {
		conv.channels[id] = make(chan string, conv.size)
	}
}

func (conv *ConveyerImpl) getChannel(id string) (chan string, error) {
	ch, exists := conv.channels[id]
	if !exists {
		return nil, ErrChan
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
	errCh := make(chan error, 1)

	for _, dec := range conv.decorators {
		conv.createChannel(dec.input)
		conv.createChannel(dec.output)
	}

	for _, mul := range conv.multiplexers {
		for _, in := range mul.inputs {
			conv.createChannel(in)
		}

		conv.createChannel(mul.output)
	}

	for _, sep := range conv.separators {
		conv.createChannel(sep.input)

		for _, out := range sep.outputs {
			conv.createChannel(out)
		}
	}

	for _, dec := range conv.decorators {
		go func(spec decoratorSpec) {
			inCh, _ := conv.getChannel(spec.input)
			outCh, _ := conv.getChannel(spec.output)

			if err := spec.fn(ctx, inCh, outCh); err != nil {
				errCh <- fmt.Errorf("decorator handler error: %v", err)
			}
		}(dec)
	}

	for _, mul := range conv.multiplexers {
		go func(spec multiplexerSpec) {
			inputs := make([]chan string, len(spec.inputs))

			for i, id := range spec.inputs {
				ch, _ := conv.getChannel(id)
				inputs[i] = ch
			}

			outCh, _ := conv.getChannel(spec.output)
			if err := spec.fn(ctx, inputs, outCh); err != nil {
				errCh <- fmt.Errorf("multiplexer handler error: %v", err)
			}
		}(mul)
	}

	for _, sep := range conv.separators {
		go func(spec separatorSpec) {
			inCh, _ := conv.getChannel(spec.input)
			outputs := make([]chan string, len(spec.outputs))
			for i, id := range spec.outputs {
				ch, _ := conv.getChannel(id)
				outputs[i] = ch
			}
			if err := spec.fn(ctx, inCh, outputs); err != nil {
				errCh <- fmt.Errorf("separator handler error: %v", err)
			}
		}(sep)
	}

	<-ctx.Done()

	for _, ch := range conv.channels {
		close(ch)
	}

	return errors.New(ctx.Err().Error())
}

func (conv *ConveyerImpl) Send(input string, data string) error {
	ch, err := conv.getChannel(input)
	if err != nil {
		return err
	}

	ch <- data

	return nil
}

func (conv *ConveyerImpl) Recv(output string) (string, error) {
	ch, err := conv.getChannel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return data, nil
}
