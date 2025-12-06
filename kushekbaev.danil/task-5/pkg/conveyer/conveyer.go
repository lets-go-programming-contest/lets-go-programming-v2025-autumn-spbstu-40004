package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

const undefined = "undefined"

var ErrChannelNotFound = errors.New("chan not found")

type Conveyer struct {
	size     int
	channels map[string]chan string
	handlers []func(context.Context) error
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]func(context.Context) error, 0),
	}
}

func (cnv *Conveyer) makeChannel(name string) {
	if _, ok := cnv.channels[name]; !ok {
		cnv.channels[name] = make(chan string, cnv.size)
	}
}

func (cnv *Conveyer) makeChannels(names ...string) {
	for _, n := range names {
		cnv.makeChannel(n)
	}
}

func (cnv *Conveyer) RegisterDecorator(
	handler func(ctx context.Context, input chan string, output chan string) error,
	input string,
	out string,
) {
	cnv.makeChannels(input)
	cnv.makeChannels(out)
	cnv.handlers = append(cnv.handlers, func(ctx context.Context) error {
		return handler(ctx, cnv.channels[input], cnv.channels[out])
	})
}

func (cnv *Conveyer) RegisterMultiplexer(
	handler func(ctx context.Context, inputs []chan string, output chan string) error,
	inNames []string,
	out string,
) {
	cnv.makeChannels(inNames...)
	cnv.makeChannels(out)
	cnv.handlers = append(cnv.handlers, func(ctx context.Context) error {
		inputChans := make([]chan string, 0, len(inNames))
		for _, name := range inNames {
			inputChans = append(inputChans, cnv.channels[name])
		}

		return handler(ctx, inputChans, cnv.channels[out])
	})
}

func (cnv *Conveyer) RegisterSeparator(
	handler func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outNames []string,
) {
	cnv.makeChannels(input)
	cnv.makeChannels(outNames...)
	cnv.handlers = append(cnv.handlers, func(ctx context.Context) error {
		outputChans := make([]chan string, 0, len(outNames))
		for _, name := range outNames {
			outputChans = append(outputChans, cnv.channels[name])
		}

		return handler(ctx, cnv.channels[input], outputChans)
	})
}

func (cnv *Conveyer) Run(ctx context.Context) error {
	errGroup, egCtx := errgroup.WithContext(ctx)

	for _, handlerFunc := range cnv.handlers {
		hf := handlerFunc

		errGroup.Go(func() error {
			return hf(egCtx)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer handlers: %w", err)
	}

	return nil
}

func (cnv *Conveyer) Send(input string, data string) error {
	channel, ok := cnv.channels[input]
	if !ok {
		return ErrChannelNotFound
	}
	channel <- data

	return nil
}

func (cnv *Conveyer) Recv(output string) (string, error) {
	channel, ok := cnv.channels[output]

	if !ok {
		return "", ErrChannelNotFound
	}

	value, open := <-channel

	if !open {
		return undefined, nil
	}

	return value, nil
}
