package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrTimeout      = errors.New("timeout")
	ErrFullChannel  = errors.New("channel is full")
)

const (
	undefinedStr = "undefined"
	timeoutTime  = 100
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
	mu           sync.RWMutex
	closeOnce    sync.Once
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

func (c *DefaultConveyer) RegisterDecorator(
	decoratorFunc func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.obtainChannel(input)
	c.obtainChannel(output)

	c.mu.Lock()
	c.decorators = append(c.decorators, specDecorator{
		fn:     decoratorFunc,
		input:  input,
		output: output,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) RegisterMultiplexer(
	multiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inputName := range inputs {
		c.obtainChannel(inputName)
	}

	c.obtainChannel(output)

	c.mu.Lock()
	c.multiplexers = append(c.multiplexers, specMultiplexer{
		fn:     multiplexerFunc,
		inputs: inputs,
		output: output,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) RegisterSeparator(
	separatorFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.obtainChannel(input)

	for _, outputName := range outputs {
		c.obtainChannel(outputName)
	}

	c.mu.Lock()
	c.separators = append(c.separators, specSeparator{
		fn:      separatorFunc,
		input:   input,
		outputs: outputs,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	errorGroup, groupCtx := errgroup.WithContext(ctx)

	launchHandler := func(handler func() error) {
		errorGroup.Go(handler)
	}

	c.mu.RLock()
	decoratorsCopy := append([]specDecorator(nil), c.decorators...)
	multiplexersCopy := append([]specMultiplexer(nil), c.multiplexers...)
	separatorsCopy := append([]specSeparator(nil), c.separators...)
	c.mu.RUnlock()

	for _, decorator := range decoratorsCopy {
		current := decorator

		launchHandler(func() error {
			input, err := c.getChannel(current.input)
			if err != nil {
				return fmt.Errorf("decorator: %w", err)
			}
			output, err := c.getChannel(current.output)
			if err != nil {
				return fmt.Errorf("decorator: %w", err)
			}

			return current.fn(groupCtx, input, output)
		})
	}

	for _, multiplexer := range multiplexersCopy {
		current := multiplexer

		launchHandler(func() error {
			inputs := make([]chan string, len(current.inputs))
			for i, name := range current.inputs {
				ch, err := c.getChannel(name)
				if err != nil {
					return fmt.Errorf("multiplexer input %q: %w", name, err)
				}
				inputs[i] = ch
			}

			output, err := c.getChannel(current.output)
			if err != nil {
				return fmt.Errorf("multiplexer output %q: %w", current.output, err)
			}

			return current.fn(groupCtx, inputs, output)
		})
	}

	for _, separator := range separatorsCopy {
		current := separator

		launchHandler(func() error {
			input, err := c.getChannel(current.input)
			if err != nil {
				return fmt.Errorf("separator input %q: %w", current.input, err)
			}

			outputs := make([]chan string, len(current.outputs))
			for i, name := range current.outputs {
				ch, err := c.getChannel(name)
				if err != nil {
					return fmt.Errorf("separator output %q: %w", name, err)
				}
				outputs[i] = ch
			}

			return current.fn(groupCtx, input, outputs)
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *DefaultConveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case channel <- data:
		return nil
	case <-time.After(timeoutTime * time.Millisecond):
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
		return undefinedStr, nil
	}

	return data, nil
}

func (c *DefaultConveyer) closeAllChannels() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		for name, ch := range c.channels {
			if ch != nil {
				close(ch)
			}
			delete(c.channels, name)
		}
	})
}
