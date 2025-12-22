package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const undefined = "undefined"

var ErrChannelNotFound = errors.New("chan not found")

type Conveyer struct {
	size     int
	channels map[string]chan string
	handlers []func(context.Context) error
	mu       sync.RWMutex
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]func(context.Context) error, 0),
		mu:       sync.RWMutex{},
	}
}

func (cnv *Conveyer) makeChannel(name string) {
	cnv.mu.Lock()
	defer cnv.mu.Unlock()

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
	cnv.makeChannels(input, out)

	cnv.mu.Lock()
	cnv.handlers = append(cnv.handlers, func(ctx context.Context) error {
		cnv.mu.RLock()
		inputChan := cnv.channels[input]
		outputChan := cnv.channels[out]
		cnv.mu.RUnlock()

		return handler(ctx, inputChan, outputChan)
	})
	cnv.mu.Unlock()
}

func (cnv *Conveyer) RegisterMultiplexer(
	handler func(ctx context.Context, inputs []chan string, output chan string) error,
	inNames []string,
	out string,
) {
	cnv.makeChannels(append(inNames, out)...)

	cnv.mu.Lock()
	cnv.handlers = append(cnv.handlers, func(ctx context.Context) error {
		cnv.mu.RLock()
		inputChans := make([]chan string, 0, len(inNames))

		for _, name := range inNames {
			inputChans = append(inputChans, cnv.channels[name])
		}

		outputChan := cnv.channels[out]
		cnv.mu.RUnlock()

		return handler(ctx, inputChans, outputChan)
	})
	cnv.mu.Unlock()
}

func (cnv *Conveyer) RegisterSeparator(
	handler func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outNames []string,
) {
	cnv.makeChannels(append([]string{input}, outNames...)...)

	cnv.mu.Lock()
	cnv.handlers = append(cnv.handlers, func(ctx context.Context) error {
		cnv.mu.RLock()
		inputChan := cnv.channels[input]
		outputChans := make([]chan string, 0, len(outNames))

		for _, name := range outNames {
			outputChans = append(outputChans, cnv.channels[name])
		}

		cnv.mu.RUnlock()

		return handler(ctx, inputChan, outputChans)
	})
	cnv.mu.Unlock()
}

func (cnv *Conveyer) Run(ctx context.Context) error {
	cnv.mu.RLock()
	handlers := make([]func(context.Context) error, len(cnv.handlers))
	copy(handlers, cnv.handlers)
	cnv.mu.RUnlock()

	errGroup, egCtx := errgroup.WithContext(ctx)

	for _, handlerFunc := range handlers {
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
	cnv.mu.RLock()
	channel, ok := cnv.channels[input]
	cnv.mu.RUnlock()

	if !ok {
		return ErrChannelNotFound
	}

	channel <- data

	return nil
}

func (cnv *Conveyer) Recv(output string) (string, error) {
	cnv.mu.RLock()
	channel, ok := cnv.channels[output]
	cnv.mu.RUnlock()

	if !ok {
		return "", ErrChannelNotFound
	}

	value, open := <-channel

	if !open {
		return undefined, nil
	}

	return value, nil
}

func (cnv *Conveyer) GetChannel(name string) (chan string, error) {
	cnv.mu.RLock()
	defer cnv.mu.RUnlock()

	channel, ok := cnv.channels[name]
	if !ok {
		return nil, ErrChannelNotFound
	}

	return channel, nil
}

func (cnv *Conveyer) HasChannel(name string) bool {
	cnv.mu.RLock()
	defer cnv.mu.RUnlock()

	_, ok := cnv.channels[name]

	return ok
}

func (cnv *Conveyer) CloseAllChannels() {
	cnv.mu.Lock()
	defer cnv.mu.Unlock()

	for name, ch := range cnv.channels {
		close(ch)
		delete(cnv.channels, name)
	}
}

func (cnv *Conveyer) CloseChannel(name string) error {
	cnv.mu.Lock()
	defer cnv.mu.Unlock()

	channel, ok := cnv.channels[name]
	if !ok {
		return ErrChannelNotFound
	}

	close(channel)
	delete(cnv.channels, name)

	return nil
}
