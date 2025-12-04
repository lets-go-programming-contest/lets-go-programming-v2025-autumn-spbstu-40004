package conveyer

import (
	"context"
	"errors"
	"sync"
)

var errChannelNotFound = errors.New("chan not found")

type Conveyer struct {
	channels    map[string]chan string
	channelSize int
	handlers    []func(ctx context.Context) error
	mutex       sync.RWMutex
}

func New(size int) Conveyer {
	return Conveyer{
		channels:    make(map[string]chan string),
		channelSize: size,
		handlers:    make([]func(ctx context.Context) error, 0),
		mutex:       sync.RWMutex{},
	}
}

func (c *Conveyer) addChannel(channelName string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.channels[channelName]; !ok {
		c.channels[channelName] = make(chan string, c.channelSize)
	}
}

func (c *Conveyer) getChannel(channelName string) (chan string, error) {
	if channel, ok := c.channels[channelName]; ok {
		return channel, nil
	}

	return nil, errChannelNotFound
}

func (c *Conveyer) closeAllChannels() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, channel := range c.channels {
		close(channel)
	}
}

func (c *Conveyer) addHandler(function func(ctx context.Context) error) {
	c.handlers = append(c.handlers, function)
}

func (c *Conveyer) RegisterDecorator(
	fn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string, output string,
) {
	c.addChannel(input)
	c.addChannel(output)

	c.addHandler(func(ctx context.Context) error {
		c.mutex.RLock()
		defer c.mutex.RUnlock()

		inputChannel, _ := c.getChannel(input)

		outputChannel, _ := c.getChannel(output)

		return fn(ctx, inputChannel, outputChannel)
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string, output string,
) {
	for i := 0; i < len(inputs); i++ {
		c.addChannel(inputs[i])
	}

	c.addChannel(output)

	c.addHandler(func(ctx context.Context) error {
		c.mutex.RLock()
		defer c.mutex.RUnlock()

		inputChannels := make([]chan string, len(inputs))
		for i := 0; i < len(inputs); i++ {
			currentChannel, _ := c.getChannel(inputs[i])

			inputChannels[i] = currentChannel
		}

		outputChannel, _ := c.getChannel(output)

		return fn(ctx, inputChannels, outputChannel)
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string, outputs []string,
) {
	c.addChannel(input)

	for i := 0; i < len(outputs); i++ {
		c.addChannel(outputs[i])
	}

	c.addHandler(func(ctx context.Context) error {
		c.mutex.RLock()
		defer c.mutex.RUnlock()

		inputChannel, _ := c.getChannel(input)

		outputChannels := make([]chan string, len(outputs))
		for i := 0; i < len(outputs); i++ {
			outputChannel, _ := c.getChannel(outputs[i])
			outputChannels[i] = outputChannel
		}

		return fn(ctx, inputChannel, outputChannels)
	})
}
