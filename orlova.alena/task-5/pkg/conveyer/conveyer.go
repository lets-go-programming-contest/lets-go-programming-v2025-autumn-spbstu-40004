package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChan = errors.New("chan not found")
)

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

	group, groupCtx := errgroup.WithContext(ctx)

	for _, decorator := range conv.decorators {
		dec := decorator

		group.Go(func() error {
			input, _ := conv.getChannel(dec.input)
			output, _ := conv.getChannel(dec.output)

			return dec.fn(groupCtx, input, output)
		})
	}

	for _, multiplexer := range conv.multiplexers {
		mul := multiplexer

		group.Go(func() error {
			inputs := make([]chan string, len(mul.inputs))

			for i, name := range mul.inputs {
				inputs[i], _ = conv.getChannel(name)
			}

			output, _ := conv.getChannel(mul.output)

			return mul.fn(groupCtx, inputs, output)
		})
	}

	for _, separator := range conv.separators {
		sep := separator

		group.Go(func() error {
			input, _ := conv.getChannel(sep.input)
			outputs := make([]chan string, len(sep.outputs))

			for i, name := range sep.outputs {
				outputs[i], _ = conv.getChannel(name)
			}

			return sep.fn(groupCtx, input, outputs)
		})
	}

	err := group.Wait()

	for _, ch := range conv.channels {
		close(ch)
	}

	if err != nil {
		return fmt.Errorf("conveyer error: %w", err)
	}

	return nil
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
