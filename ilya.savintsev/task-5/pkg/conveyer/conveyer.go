package conveyer

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrTimeout      = errors.New("timeout")
	ErrFullChannel  = errors.New("channel is full")
)

type specDecorator struct {
	fn     func(ctx context.Context, input chan string, output chan string) error
	input  string
	output string
}

type specMultiplexer struct {
	fn     func(ctx context.Context, inputs []chan string, output chan string) error
	inputs []string
	output string
}

type specSeparator struct {
	fn      func(ctx context.Context, input chan string, outputs []chan string) error
	input   string
	outputs []string
}

type DefaultConveyer struct {
	channels     map[string]chan string
	bufferSize   int
	decorators   []specDecorator
	multiplexers []specMultiplexer
	separators   []specSeparator
}

func New(size int) *DefaultConveyer {
	return &DefaultConveyer{
		channels:     make(map[string]chan string),
		bufferSize:   size,
		decorators:   make([]specDecorator, 0),
		multiplexers: make([]specMultiplexer, 0),
		separators:   make([]specSeparator, 0),
	}
}

func (c *DefaultConveyer) getOrCreateChannel(name string) chan string {
	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.bufferSize)
	c.channels[name] = channel

	return channel
}

func (c *DefaultConveyer) getChannel(name string) (chan string, error) {
	if channel, exists := c.channels[name]; exists {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *DefaultConveyer) RegisterDecorator(
	decoratorFunc func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.decorators = append(c.decorators, specDecorator{
		fn:     decoratorFunc,
		input:  input,
		output: output,
	})
}

func (c *DefaultConveyer) RegisterMultiplexer(
	multiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inputName := range inputs {
		c.getOrCreateChannel(inputName)
	}

	c.getOrCreateChannel(output)

	c.multiplexers = append(c.multiplexers, specMultiplexer{
		fn:     multiplexerFunc,
		inputs: inputs,
		output: output,
	})
}

func (c *DefaultConveyer) RegisterSeparator(
	separatorFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)

	for _, outputName := range outputs {
		c.getOrCreateChannel(outputName)
	}

	c.separators = append(c.separators, specSeparator{
		fn:      separatorFunc,
		input:   input,
		outputs: outputs,
	})
}

func (c *DefaultConveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	errorGroup, groupCtx := errgroup.WithContext(ctx)

	launchHandler := func(handler func() error) {
		errorGroup.Go(handler)
	}

	for _, decorator := range c.decorators {
		current := decorator
		launchHandler(func() error {
			input, _ := c.getChannel(current.input)
			output, _ := c.getChannel(current.output)
			return current.fn(groupCtx, input, output)
		})
	}

	for _, multiplexer := range c.multiplexers {
		current := multiplexer
		launchHandler(func() error {
			inputs := make([]chan string, len(current.inputs))
			for i, name := range current.inputs {
				inputs[i], _ = c.getChannel(name)
			}
			output, _ := c.getChannel(current.output)
			return current.fn(groupCtx, inputs, output)
		})
	}

	for _, separator := range c.separators {
		current := separator
		launchHandler(func() error {
			input, _ := c.getChannel(current.input)
			outputs := make([]chan string, len(current.outputs))
			for i, name := range current.outputs {
				outputs[i], _ = c.getChannel(name)
			}
			return current.fn(groupCtx, input, outputs)
		})
	}

	return errorGroup.Wait()
}

func (c *DefaultConveyer) closeAllChannels() {
	for _, channel := range c.channels {
		close(channel)
	}
}

func (c *DefaultConveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case channel <- data:
		return nil
	case <-time.After(100 * time.Millisecond):
		return ErrTimeout
	default:
		return ErrFullChannel
	}
}

func (c *DefaultConveyer) Recv(output string) (string, error) {
	channel, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-channel
	if !ok {
		return "undefined", nil
	}

	return data, nil
}
