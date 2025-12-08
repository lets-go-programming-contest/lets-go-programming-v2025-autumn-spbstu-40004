package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var errChannelNotFound = errors.New("chan not found")

const undefined = "undefined"

type ConveyerInterface interface {
	RegisterDecorator(
		function func(
			ctx context.Context,
			input chan string,
			output chan string,
		) error,
		input string,
		output string,
	)

	RegisterMultiplexer(
		function func(
			ctx context.Context,
			inputs []chan string,
			output chan string,
		) error,
		inputs []string,
		output string,
	)

	RegisterSeparator(
		function func(
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
	function func(
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

		return function(ctx, inputChannel, outputChannel)
	})
}

func (c *Conveyer) RegisterMultiplexer(
	function func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string, output string,
) {
	for _, input := range inputs {
		c.addChannel(input)
	}

	c.addChannel(output)

	c.addHandler(func(ctx context.Context) error {
		c.mutex.RLock()
		defer c.mutex.RUnlock()

		inputChannels := make([]chan string, len(inputs))

		for index, input := range inputs {
			currentChannel, _ := c.getChannel(input)

			inputChannels[index] = currentChannel
		}

		outputChannel, _ := c.getChannel(output)

		return function(ctx, inputChannels, outputChannel)
	})
}

func (c *Conveyer) RegisterSeparator(
	function func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string, outputs []string,
) {
	c.addChannel(input)

	for _, output := range outputs {
		c.addChannel(output)
	}

	c.addHandler(func(ctx context.Context) error {
		c.mutex.RLock()
		defer c.mutex.RUnlock()

		inputChannel, _ := c.getChannel(input)

		outputChannels := make([]chan string, len(outputs))

		for index, output := range outputs {
			outputChannel, _ := c.getChannel(output)
			outputChannels[index] = outputChannel
		}

		return function(ctx, inputChannel, outputChannels)
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	errGroup, curCtx := errgroup.WithContext(ctx)

	c.mutex.RLock()

	for _, handler := range c.handlers {
		errGroup.Go(func() error {
			return handler(curCtx)
		})
	}

	c.mutex.RUnlock()

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("run function failed: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	c.mutex.RLock()

	inputChannel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	c.mutex.RUnlock()

	inputChannel <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	c.mutex.RLock()

	outputChannel, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	c.mutex.RUnlock()

	if data, ok := <-outputChannel; ok {
		return data, nil
	} else {
		return undefined, nil
	}
}
