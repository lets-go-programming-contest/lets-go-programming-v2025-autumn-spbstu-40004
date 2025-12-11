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
	fn func(context.Context, chan string, chan string) error,
	input, output string,
) {
	c.obtainChannel(input)
	c.obtainChannel(output)

	c.mu.Lock()
	c.decorators = append(c.decorators, specDecorator{
		fn:     fn,
		input:  input,
		output: output,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	for _, name := range inputs {
		c.obtainChannel(name)
	}

	c.obtainChannel(output)

	c.mu.Lock()
	c.multiplexers = append(c.multiplexers, specMultiplexer{
		fn:     fn,
		inputs: inputs,
		output: output,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.obtainChannel(input)

	for _, name := range outputs {
		c.obtainChannel(name)
	}

	c.mu.Lock()
	c.separators = append(c.separators, specSeparator{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
	c.mu.Unlock()
}

func (c *DefaultConveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	g, gctx := errgroup.WithContext(ctx)

	c.runDecorators(g, gctx)
	c.runMultiplexers(g, gctx)
	c.runSeparators(g, gctx)

	err := g.Wait()
	if err != nil {
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

func (c *DefaultConveyer) runDecorators(g *errgroup.Group, ctx context.Context) {
	c.mu.RLock()
	copyList := append([]specDecorator(nil), c.decorators...)
	c.mu.RUnlock()

	for _, d := range copyList {
		dec := d

		g.Go(func() error {
			in, err := c.getChannel(dec.input)
			if err != nil {
				return err
			}

			out, err := c.getChannel(dec.output)
			if err != nil {
				return err
			}

			return dec.fn(ctx, in, out)
		})
	}
}

func (c *DefaultConveyer) runMultiplexers(g *errgroup.Group, ctx context.Context) {
	c.mu.RLock()
	copyList := append([]specMultiplexer(nil), c.multiplexers...)
	c.mu.RUnlock()

	for _, m := range copyList {
		mp := m

		g.Go(func() error {
			var inputs []chan string
			for _, name := range mp.inputs {
				channel, err := c.getChannel(name)
				if err != nil {
					return err
				}

				inputs = append(inputs, channel)
			}

			out, err := c.getChannel(mp.output)
			if err != nil {
				return err
			}

			return mp.fn(ctx, inputs, out)
		})
	}
}

func (c *DefaultConveyer) runSeparators(g *errgroup.Group, ctx context.Context) {
	c.mu.RLock()
	copyList := append([]specSeparator(nil), c.separators...)
	c.mu.RUnlock()

	for _, s := range copyList {
		sep := s

		g.Go(func() error {
			in, err := c.getChannel(sep.input)
			if err != nil {
				return err
			}

			var outputs []chan string
			for _, name := range sep.outputs {
				out, err := c.getChannel(name)
				if err != nil {
					return err
				}

				outputs = append(outputs, out)
			}

			return sep.fn(ctx, in, outputs)
		})
	}
}
