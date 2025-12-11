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
	fn     func(context.Context, chan string, chan string) error
	input  string
	output string
}

type specMultiplexer struct {
	fn     func(context.Context, []chan string, chan string) error
	inputs []string
	output string
}

type specSeparator struct {
	fn      func(context.Context, chan string, []chan string) error
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
		mu:           sync.RWMutex{},
		closeOnce:    sync.Once{},
		channels:     make(map[string]chan string),
		bufferSize:   size,
		decorators:   []specDecorator{},
		multiplexers: []specMultiplexer{},
		separators:   []specSeparator{},
	}
}

func (c *DefaultConveyer) RegisterDecorator(
	handlerFunc func(context.Context, chan string, chan string) error,
	input, output string,
) {
	c.obtainChannel(input)
	c.obtainChannel(output)

	c.mu.Lock()
	c.decorators = append(c.decorators, specDecorator{
		fn:     handlerFunc,
		input:  input,
		output: output,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) RegisterMultiplexer(
	handlerFunc func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	for _, name := range inputs {
		c.obtainChannel(name)
	}

	c.obtainChannel(output)

	c.mu.Lock()
	c.multiplexers = append(c.multiplexers, specMultiplexer{
		fn:     handlerFunc,
		inputs: inputs,
		output: output,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) RegisterSeparator(
	handlerFunc func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.obtainChannel(input)

	for _, name := range outputs {
		c.obtainChannel(name)
	}

	c.mu.Lock()
	c.separators = append(c.separators, specSeparator{
		fn:      handlerFunc,
		input:   input,
		outputs: outputs,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	group, groupCtx := errgroup.WithContext(ctx)

	c.runDecorators(group, groupCtx)
	c.runMultiplexers(group, groupCtx)
	c.runSeparators(group, groupCtx)

	if err := group.Wait(); err != nil {
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

func (c *DefaultConveyer) runDecorators(group *errgroup.Group, ctx context.Context) {
	c.mu.RLock()
	decoratorList := append([]specDecorator(nil), c.decorators...)
	c.mu.RUnlock()

	for _, decoratorSpec := range decoratorList {
		current := decoratorSpec

		group.Go(func() error {
			inputChannel, err := c.getChannel(current.input)
			if err != nil {
				return err
			}

			outputChannel, err := c.getChannel(current.output)
			if err != nil {
				return err
			}

			return current.fn(ctx, inputChannel, outputChannel)
		})
	}
}

func (c *DefaultConveyer) runMultiplexers(group *errgroup.Group, ctx context.Context) {
	c.mu.RLock()
	multiplexerList := append([]specMultiplexer(nil), c.multiplexers...)
	c.mu.RUnlock()

	for _, multiplexer := range multiplexerList {
		current := multiplexer

		group.Go(func() error {
			inputChannels := make([]chan string, 0, len(current.inputs))

			for _, name := range current.inputs {
				channel, err := c.getChannel(name)
				if err != nil {
					return err
				}

				inputChannels = append(inputChannels, channel)
			}

			outputChannel, err := c.getChannel(current.output)
			if err != nil {
				return err
			}

			return current.fn(ctx, inputChannels, outputChannel)
		})
	}
}

func (c *DefaultConveyer) runSeparators(group *errgroup.Group, ctx context.Context) {
	c.mu.RLock()
	separatorList := append([]specSeparator(nil), c.separators...)
	c.mu.RUnlock()

	for _, separator := range separatorList {
		current := separator

		group.Go(func() error {
			inputChannel, err := c.getChannel(current.input)
			if err != nil {
				return err
			}

			outputChannels := make([]chan string, 0, len(current.outputs))

			for _, name := range current.outputs {
				channel, err := c.getChannel(name)
				if err != nil {
					return err
				}

				outputChannels = append(outputChannels, channel)
			}

			return current.fn(ctx, inputChannel, outputChannels)
		})
	}
}
