package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var ErrNoChannels = errors.New("chan not found")

const errUndefined = "undefined"

type IConveyer interface {
	RegisterDecorator(
		fn func(
			ctx context.Context,
			input chan string,
			output chan string,
		) error,
		input string,
		output string,
	)
	RegisterMultiplexer(
		fn func(
			ctx context.Context,
			inputs []chan string,
			output chan string,
		) error,
		inputs []string,
		output string,
	)
	RegisterSeparator(
		fn func(
			ctx context.Context,
			input chan string,
			outputs []chan string,
		) error,
		input string,
		outputs []string,
	)
	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

type Conveyer struct {
	size     int
	channels map[string]chan string
	handlers []func(ctx context.Context) error
}

func New(size int) Conveyer {
	return Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		handlers: []func(ctx context.Context) error{},
	}
}

func (conveyer *Conveyer) CloseChannels() {
	for _, channel := range conveyer.channels {
		close(channel)
	}
}

func (conveyer *Conveyer) getOrCreateChannel(name string) chan string {
	if conveyer.channels == nil {
		conveyer.channels = make(map[string]chan string)
	}

	handlerChannel, exist := conveyer.channels[name]
	if !exist {
		handlerChannel = make(chan string, conveyer.size)
		conveyer.channels[name] = handlerChannel
	}

	return handlerChannel
}

func (conveyer *Conveyer) RegisterDecorator(
	decorator func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	inputChan := conveyer.getOrCreateChannel(input)
	outputChan := conveyer.getOrCreateChannel(output)

	handler := func(ctx context.Context) error {
		return decorator(ctx, inputChan, outputChan)
	}

	conveyer.handlers = append(conveyer.handlers, handler)
}

func (conveyer *Conveyer) RegisterMultiplexer(
	multiplexer func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	inputChans := make([]chan string, len(inputs))
	for i, name := range inputs {
		inputChans[i] = conveyer.getOrCreateChannel(name)
	}

	outputChan := conveyer.getOrCreateChannel(output)

	handler := func(ctx context.Context) error {
		return multiplexer(ctx, inputChans, outputChan)
	}

	conveyer.handlers = append(conveyer.handlers, handler)
}

func (conveyer *Conveyer) RegisterSeparator(
	separator func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	inputChan := conveyer.getOrCreateChannel(input)

	outputChans := make([]chan string, len(outputs))
	for i, name := range outputs {
		outputChans[i] = conveyer.getOrCreateChannel(name)
	}

	handler := func(ctx context.Context) error {
		return separator(ctx, inputChan, outputChans)
	}

	conveyer.handlers = append(conveyer.handlers, handler)
}

func (conveyer *Conveyer) Run(ctx context.Context) error {
	if len(conveyer.handlers) == 0 {
		return nil
	}

	defer conveyer.CloseChannels()

	handlerErrGroup, groupCtx := errgroup.WithContext(ctx)

	for _, handler := range conveyer.handlers {
		handlerErrGroup.Go(func() error {
			return handler(groupCtx)
		})
	}

	if err := handlerErrGroup.Wait(); err != nil {
		return fmt.Errorf("handler error: %w", err)
	}

	return nil
}

func (conveyer *Conveyer) Send(input string, data string) error {
	channel, exist := conveyer.channels[input]
	if !exist {
		return ErrNoChannels
	}
	channel <- data

	return nil
}

func (conveyer *Conveyer) Recv(output string) (string, error) {
	channel, exist := conveyer.channels[output]
	if !exist {
		return "", ErrNoChannels
	}

	str, ok := <-channel
	if !ok {
		return errUndefined, nil
	}

	return str, nil
}
