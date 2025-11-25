package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrUndefined    = "undefined"
)

type conveyer struct {
	mu          sync.RWMutex
	channels    map[string]chan string
	size        int
	decorators  []decoratorConfig
	multiplexers []multiplexerConfig
	separators  []separatorConfig
	cancel      context.CancelFunc
	wg          sync.WaitGroup
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
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

func (c *conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

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
	inputChan := c.getOrCreateChannel(input)
	outputChan := c.getOrCreateChannel(output)

	c.decorators = append(c.decorators, decoratorConfig{
		fn:     fn,
		input:  input,
		output: output,
	})

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
	}()
}

func (c *conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inputChans := make([]chan string, len(inputs))
	for i, input := range inputs {
		inputChans[i] = c.getOrCreateChannel(input)
	}
	outputChan := c.getOrCreateChannel(output)

	c.multiplexers = append(c.multiplexers, multiplexerConfig{
		fn:     fn,
		inputs: inputs,
		output: output,
	})

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
	}()
}

func (c *conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inputChan := c.getOrCreateChannel(input)
	outputChans := make([]chan string, len(outputs))
	for i, output := range outputs {
		outputChans[i] = c.getOrCreateChannel(output)
	}

	c.separators = append(c.separators, separatorConfig{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
	}()
}

func (c *conveyer) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	c.cancel = cancel
	defer cancel()

	for _, decorator := range c.decorators {
		inputChan, _ := c.getChannel(decorator.input)
		outputChan, _ := c.getChannel(decorator.output)

		c.wg.Add(1)
		go func(d decoratorConfig, in, out chan string) {
			defer c.wg.Done()
			if err := d.fn(ctx, in, out); err != nil {
				cancel()
			}
		}(decorator, inputChan, outputChan)
	}

	for _, multiplexer := range c.multiplexers {
		inputChans := make([]chan string, len(multiplexer.inputs))
		for i, input := range multiplexer.inputs {
			inputChans[i], _ = c.getChannel(input)
		}
		outputChan, _ := c.getChannel(multiplexer.output)

		c.wg.Add(1)
		go func(m multiplexerConfig, in []chan string, out chan string) {
			defer c.wg.Done()
			if err := m.fn(ctx, in, out); err != nil {
				cancel()
			}
		}(multiplexer, inputChans, outputChan)
	}

	for _, separator := range c.separators {
		inputChan, _ := c.getChannel(separator.input)
		outputChans := make([]chan string, len(separator.outputs))
		for i, output := range separator.outputs {
			outputChans[i], _ = c.getChannel(output)
		}

		c.wg.Add(1)
		go func(s separatorConfig, in chan string, outs []chan string) {
			defer c.wg.Done()
			if err := s.fn(ctx, in, outs); err != nil {
				cancel()
			}
		}(separator, inputChan, outputChans)
	}

	select {
	case <-ctx.Done():
		c.closeAllChannels()
		c.wg.Wait()
		return ctx.Err()
	default:
		c.wg.Wait()
		c.closeAllChannels()
		return nil
	}
}

func (c *conveyer) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}
}

func (c *conveyer) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	select {
	case data, ok := <-ch:
		if !ok {
			return ErrUndefined, nil
		}
		return data, nil
	default:
		return "", errors.New("no data available")
	}
}
