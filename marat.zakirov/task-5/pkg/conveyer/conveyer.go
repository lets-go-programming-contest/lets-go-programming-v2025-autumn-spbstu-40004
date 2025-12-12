package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrChannel = errors.New("chan not found")

const ClosedInChannel = "undefined"

type Conveyer struct {
	channels map[string]chan string
	handlers []func(ctx context.Context) error
	size     int
	mu       sync.RWMutex
}

func New(size int) Conveyer {
	return Conveyer{
		channels: make(map[string]chan string),
		handlers: make([]func(ctx context.Context) error, 0),
		size:     size,
		mu:       sync.RWMutex{},
	}
}

func (c *Conveyer) createChannel(chName string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if channel, exists := c.channels[chName]; exists {
		return channel
	}

	newChannel := make(chan string, c.size)
	c.channels[chName] = newChannel

	return newChannel
}

func (c *Conveyer) closeChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		close(channel)
	}
}

func (c *Conveyer) RegisterDecorator(
	function func(cntx context.Context, inChannelP chan string, outChannelP chan string) error,
	inChannelP string,
	outChannelP string,
) {
	inChannel := c.createChannel(inChannelP)
	outChannel := c.createChannel(outChannelP)

	c.mu.Lock()
	c.handlers = append(c.handlers, func(cntx context.Context) error {
		return function(cntx, inChannel, outChannel)
	})
	c.mu.Unlock()
}

func (c *Conveyer) RegisterMultiplexer(
	function func(ctx context.Context, inChannelsP []chan string, outChannelP chan string) error,
	inChannelsP []string,
	outChannelP string,
) {
	inChannels := make([]chan string, len(inChannelsP))
	for i, name := range inChannelsP {
		inChannels[i] = c.createChannel(name)
	}

	outChannel := c.createChannel(outChannelP)

	c.mu.Lock()
	c.handlers = append(c.handlers, func(cntx context.Context) error {
		return function(cntx, inChannels, outChannel)
	})
	c.mu.Unlock()
}

func (c *Conveyer) RegisterSeparator(
	function func(ctx context.Context, inChannelP chan string, outChannelsP []chan string) error,
	inChannelP string,
	outChannelsP []string,
) {
	inChannel := c.createChannel(inChannelP)
	outChannels := make([]chan string, len(outChannelsP))

	for i, name := range outChannelsP {
		outChannels[i] = c.createChannel(name)
	}

	c.mu.Lock()
	c.handlers = append(c.handlers, func(cntx context.Context) error {
		return function(cntx, inChannel, outChannels)
	})
	c.mu.Unlock()
}

func (c *Conveyer) Run(cntx context.Context) error {
	defer c.closeChannels()

	errorGroup, cntx := errgroup.WithContext(cntx)
	for _, fn := range c.handlers {
		errorGroup.Go(func() error {
			return fn(cntx)
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return fmt.Errorf("error handler: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(inChannelP string, stringData string) error {
	inChannel, exists := c.channels[inChannelP]
	if !exists {
		return ErrChannel
	}

	inChannel <- stringData

	return nil
}

func (c *Conveyer) Recv(outChannelP string) (string, error) {
	outChannel, exists := c.channels[outChannelP]

	if !exists {
		return "", ErrChannel
	}

	str, ok := <-outChannel
	if !ok {
		return ClosedInChannel, nil
	}

	return str, nil
}
