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
