package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrChanNotFound = errors.New("chan not found")

const errUndefined = "undefined"

type conveyer struct {
	mu       sync.RWMutex
	channels map[string]chan string
	size     int
	handlers []func(ctx context.Context) error
}

func New(size int) *conveyer {
	return &conveyer{
		channels: make(map[string]chan string),
		size:     size,
		handlers: make([]func(ctx context.Context) error, 0),
		mu:       sync.RWMutex{},
	}
}

func (c *conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.size)
	c.channels[name] = channel

	return channel
}

func (c *conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if channel, exists := c.channels[name]; exists {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *conveyer) RegisterDecorator(
	decoratorFunc func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inputChan := c.getOrCreateChannel(input)
	outputChan := c.getOrCreateChannel(output)

	handler := func(ctx context.Context) error {
		return decoratorFunc(ctx, inputChan, outputChan)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, handler)
	c.mu.Unlock()
}

func (c *conveyer) RegisterMultiplexer(
	multiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inputChans := make([]chan string, len(inputs))
	for i, name := range inputs {
		inputChans[i] = c.getOrCreateChannel(name)
	}

	outputChan := c.getOrCreateChannel(output)

	handler := func(ctx context.Context) error {
		return multiplexerFunc(ctx, inputChans, outputChan)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, handler)
	c.mu.Unlock()
}

func (c *conveyer) RegisterSeparator(
	separatorFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inputChan := c.getOrCreateChannel(input)

	outputChans := make([]chan string, len(outputs))
	for i, name := range outputs {
		outputChans[i] = c.getOrCreateChannel(name)
	}

	handler := func(ctx context.Context) error {
		return separatorFunc(ctx, inputChan, outputChans)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, handler)
	c.mu.Unlock()
}

func (c *conveyer) closeAllChannels() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, channel := range c.channels {
		close(channel)
	}
}

func (c *conveyer) Run(ctx context.Context) error {
	if len(c.handlers) == 0 {
		return nil
	}

	defer c.closeAllChannels()

	errorGroup, groupCtx := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		h := handler

		errorGroup.Go(func() error {
			return h(groupCtx)
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *conveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	channel <- data

	return nil
}

func (c *conveyer) Recv(output string) (string, error) {
	channel, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-channel
	if !ok {
		return errUndefined, nil
	}

	return data, nil
}
