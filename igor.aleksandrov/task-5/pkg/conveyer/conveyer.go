package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const UndefinedValue = "undefined"

var ErrChannelNotFound = errors.New("chan not found")

type Conveyer struct {
	mu          sync.RWMutex
	channels    map[string]chan string
	handlers    []func(ctx context.Context) error
	channelSize int
}

func New(size int) *Conveyer {
	return &Conveyer{
		mu:          sync.RWMutex{},
		channels:    make(map[string]chan string),
		handlers:    make([]func(ctx context.Context) error, 0),
		channelSize: size,
	}
}

func (c *Conveyer) ensureChannel(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.channels[name]; !exists {
		c.channels[name] = make(chan string, c.channelSize)
	}
}

func (c *Conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	channel, exists := c.channels[name]

	if !exists {
		return nil, ErrChannelNotFound
	}

	return channel, nil
}

func (c *Conveyer) getExistingChannelUnsafe(name string) chan string {
	return c.channels[name]
}

func (c *Conveyer) RegisterDecorator(
	handler func(context.Context, chan string, chan string) error,
	inputName string,
	outputName string,
) {
	c.ensureChannel(inputName)
	c.ensureChannel(outputName)

	wrapper := func(ctx context.Context) error {
		c.mu.RLock()
		inCh := c.getExistingChannelUnsafe(inputName)
		outCh := c.getExistingChannelUnsafe(outputName)
		c.mu.RUnlock()

		return handler(ctx, inCh, outCh)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, wrapper)
	c.mu.Unlock()
}

func (c *Conveyer) RegisterMultiplexer(
	handler func(context.Context, []chan string, chan string) error,
	inputs []string,
	outputName string,
) {
	for _, name := range inputs {
		c.ensureChannel(name)
	}

	c.ensureChannel(outputName)

	wrapper := func(ctx context.Context) error {
		c.mu.RLock()
		inChs := make([]chan string, 0, len(inputs))

		for _, name := range inputs {
			inChs = append(inChs, c.getExistingChannelUnsafe(name))
		}

		outCh := c.getExistingChannelUnsafe(outputName)
		c.mu.RUnlock()

		return handler(ctx, inChs, outCh)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, wrapper)
	c.mu.Unlock()
}

func (c *Conveyer) RegisterSeparator(
	handler func(context.Context, chan string, []chan string) error,
	inputName string,
	outputs []string,
) {
	c.ensureChannel(inputName)

	for _, name := range outputs {
		c.ensureChannel(name)
	}

	wrapper := func(ctx context.Context) error {
		c.mu.RLock()
		inCh := c.getExistingChannelUnsafe(inputName)
		outChs := make([]chan string, 0, len(outputs))

		for _, name := range outputs {
			outChs = append(outChs, c.getExistingChannelUnsafe(name))
		}

		c.mu.RUnlock()

		return handler(ctx, inCh, outChs)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, wrapper)
	c.mu.Unlock()
}

func (c *Conveyer) Run(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)

	c.mu.RLock()

	for _, hnd := range c.handlers {
		h := hnd

		group.Go(func() error {
			return h(gCtx)
		})
	}

	c.mu.RUnlock()

	err := group.Wait()

	c.closeAllChannels()

	if err != nil {
		return fmt.Errorf("conveyer run error: %w", err)
	}

	return nil
}

func (c *Conveyer) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		func() {
			defer func() {
				_ = recover()
			}()
			close(channel)
		}()
	}
}

func (c *Conveyer) Send(name string, data string) error {
	channel, err := c.getChannel(name)
	if err != nil {
		return err
	}

	defer func() {
		_ = recover()
	}()

	channel <- data

	return nil
}

func (c *Conveyer) Recv(name string) (string, error) {
	channel, err := c.getChannel(name)
	if err != nil {
		return "", err
	}

	val, ok := <-channel

	if !ok {
		return UndefinedValue, nil
	}

	return val, nil
}
