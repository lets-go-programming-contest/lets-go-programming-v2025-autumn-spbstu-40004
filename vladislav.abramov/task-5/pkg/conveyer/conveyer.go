package conveyer

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrUndefined    = "undefined"
)

type conveyer struct {
	channels    map[string]chan string
	size        int
	decorators  []decoratorConfig
	multiplexers []multiplexerConfig
	separators  []separatorConfig
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
		channels: make(map[string]chan string),
		size:     size,
	}
}

func (c *conveyer) getOrCreateChannel(name string) chan string {
	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

func (c *conveyer) getChannel(name string) (chan string, error) {
	if ch, exists := c.channels[name]; exists {
		return ch, nil
	}
	return nil, ErrChanNotFound
}

func (c *conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.decorators = append(c.decorators, decoratorConfig{
		fn:     fn,
		input:  input,
		output: output,
	})
}

func (c *conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, input := range inputs {
		c.getOrCreateChannel(input)
	}
	c.getOrCreateChannel(output)

	c.multiplexers = append(c.multiplexers, multiplexerConfig{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
}

func (c *conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)
	for _, output := range outputs {
		c.getOrCreateChannel(output)
	}

	c.separators = append(c.separators, separatorConfig{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
}

func (c *conveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	g, ctx := errgroup.WithContext(ctx)

	for _, decorator := range c.decorators {
		inputChan, _ := c.getChannel(decorator.input)
		outputChan, _ := c.getChannel(decorator.output)

		g.Go(func() error {
			return decorator.fn(ctx, inputChan, outputChan)
		})
	}

	for _, multiplexer := range c.multiplexers {
		inputChans := make([]chan string, len(multiplexer.inputs))
		for i, input := range multiplexer.inputs {
			inputChans[i], _ = c.getChannel(input)
		}
		outputChan, _ := c.getChannel(multiplexer.output)

		g.Go(func() error {
			return multiplexer.fn(ctx, inputChans, outputChan)
		})
	}

	for _, separator := range c.separators {
		inputChan, _ := c.getChannel(separator.input)
		outputChans := make([]chan string, len(separator.outputs))
		for i, output := range separator.outputs {
			outputChans[i], _ = c.getChannel(output)
		}

		g.Go(func() error {
			return separator.fn(ctx, inputChan, outputChans)
		})
	}

	return g.Wait()
}

func (c *conveyer) closeAllChannels() {
	for _, ch := range c.channels {
		close(ch)
	}
}

func (c *conveyer) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	ch <- data
	return nil
}

func (c *conveyer) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-ch
	if !ok {
		return ErrUndefined, nil
	}
	return data, nil
}
