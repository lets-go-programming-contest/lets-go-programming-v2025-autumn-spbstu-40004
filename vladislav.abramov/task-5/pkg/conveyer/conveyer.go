package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound = errors.New("chan not found")
)

const errUndefined = "undefined"

type conveyer struct {
	channels     map[string]chan string
	size         int
	decorators   []decoratorConfig
	multiplexers []multiplexerConfig
	separators   []separatorConfig
}

type decoratorConfig struct {
	fn     func(ctx context.Context, input chan string, output chan string) error
	input  string
	output string
}

type multiplexerConfig struct {
	fn     func(ctx context.Context, inputs []chan string, output chan string) error
	inputs []string
	output string
}

type separatorConfig struct {
	fn      func(ctx context.Context, input chan string, outputs []chan string) error
	input   string
	outputs []string
}

func New(size int) *conveyer {
	return &conveyer{
		channels:     make(map[string]chan string),
		size:         size,
		decorators:   make([]decoratorConfig, 0),
		multiplexers: make([]multiplexerConfig, 0),
		separators:   make([]separatorConfig, 0),
	}
}

func (c *conveyer) getOrCreateChannel(name string) chan string {
	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.size)
	c.channels[name] = channel

	return channel
}

func (c *conveyer) getChannel(name string) (chan string, error) {
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
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.decorators = append(c.decorators, decoratorConfig{
		fn:     decoratorFunc,
		input:  input,
		output: output,
	})
}

func (c *conveyer) RegisterMultiplexer(
	multiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inputName := range inputs {
		c.getOrCreateChannel(inputName)
	}

	c.getOrCreateChannel(output)

	c.multiplexers = append(c.multiplexers, multiplexerConfig{
		fn:     multiplexerFunc,
		inputs: inputs,
		output: output,
	})
}

func (c *conveyer) RegisterSeparator(
	separatorFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)

	for _, outputName := range outputs {
		c.getOrCreateChannel(outputName)
	}

	c.separators = append(c.separators, separatorConfig{
		fn:      separatorFunc,
		input:   input,
		outputs: outputs,
	})
}

func (c *conveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	errorGroup, groupCtx := errgroup.WithContext(ctx)

	for _, decorator := range c.decorators {
		inputChannel, _ := c.getChannel(decorator.input)
		outputChannel, _ := c.getChannel(decorator.output)

		currentDecorator := decorator

		errorGroup.Go(func() error {
			return currentDecorator.fn(groupCtx, inputChannel, outputChannel)
		})
	}

	for _, multiplexer := range c.multiplexers {
		inputChannels := make([]chan string, len(multiplexer.inputs))

		for index, inputName := range multiplexer.inputs {
			inputChannels[index], _ = c.getChannel(inputName)
		}

		outputChannel, _ := c.getChannel(multiplexer.output)

		currentMultiplexer := multiplexer

		errorGroup.Go(func() error {
			return currentMultiplexer.fn(groupCtx, inputChannels, outputChannel)
		})
	}

	for _, separator := range c.separators {
		inputChannel, _ := c.getChannel(separator.input)
		outputChannels := make([]chan string, len(separator.outputs))

		for index, outputName := range separator.outputs {
			outputChannels[index], _ = c.getChannel(outputName)
		}

		currentSeparator := separator

		errorGroup.Go(func() error {
			return currentSeparator.fn(groupCtx, inputChannel, outputChannels)
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *conveyer) closeAllChannels() {
	for _, channel := range c.channels {
		close(channel)
	}
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
