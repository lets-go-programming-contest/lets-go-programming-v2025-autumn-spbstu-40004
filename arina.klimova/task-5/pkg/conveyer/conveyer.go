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
	timeoutTime = 100 * time.Millisecond
)

type conveyer struct {
	mu           sync.RWMutex
	channels     map[string]chan string
	bufferSize   int
	decorators   []decoratorSpec
	multiplexers []multiplexerSpec
	separators   []separatorSpec
}

type decoratorSpec struct {
	fn     func(ctx context.Context, input chan string, output chan string) error
	input  string
	output string
}

type separatorSpec struct {
	fn      func(ctx context.Context, input chan string, outputs []chan string) error
	input   string
	outputs []string
}

type multiplexerSpec struct {
	fn     func(ctx context.Context, inputs []chan string, output chan string) error
	inputs []string
	output string
}

func New(size int) *conveyer {
	return &conveyer{
		mu:           sync.RWMutex{},
		channels:     make(map[string]chan string),
		bufferSize:   size,
		decorators:   make([]decoratorSpec, 0),
		multiplexers: make([]multiplexerSpec, 0),
		separators:   make([]separatorSpec, 0),
	}
}

func (c *conveyer) obtainChannel(name string) {
	if _, exists := c.channels[name]; exists {
		return
	}

	c.channels[name] = make(chan string, c.bufferSize)
}

func (c *conveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

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
	c.mu.Lock()
	defer c.mu.Unlock()

	c.obtainChannel(input)
	c.obtainChannel(output)

	c.decorators = append(c.decorators, decoratorSpec{
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
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, inputName := range inputs {
		if _, exists := c.channels[inputName]; !exists {
			c.obtainChannel(inputName)
		}
	}

	c.obtainChannel(output)

	c.multiplexers = append(c.multiplexers, multiplexerSpec{
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
	c.mu.Lock()
	defer c.mu.Unlock()

	c.obtainChannel(input)

	for _, outputName := range outputs {
		c.obtainChannel(outputName)
	}

	c.separators = append(c.separators, separatorSpec{
		fn:      separatorFunc,
		input:   input,
		outputs: outputs,
	})
}

func (c *conveyer) Run(ctx context.Context) error {
	decoratorsCopy, multiplexersCopy, separatorsCopy := c.createSpecCopies()

	group, groupCtx := errgroup.WithContext(ctx)

	c.runDecoratorHandlers(group, groupCtx, decoratorsCopy)
	c.runMultiplexerHandlers(group, groupCtx, multiplexersCopy)
	c.runSeparatorHandlers(group, groupCtx, separatorsCopy)

	err := group.Wait()

	c.mu.Lock()
	defer c.mu.Unlock()

	for _, ch := range c.channels {
		close(ch)
	}

	if err != nil {
		return fmt.Errorf("conveyer finished: %w", err)
	}

	return nil
}

func (c *conveyer) createSpecCopies() ([]decoratorSpec, []multiplexerSpec, []separatorSpec) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	decoratorsCopy := make([]decoratorSpec, len(c.decorators))
	copy(decoratorsCopy, c.decorators)

	multiplexersCopy := make([]multiplexerSpec, len(c.multiplexers))
	copy(multiplexersCopy, c.multiplexers)

	separatorsCopy := make([]separatorSpec, len(c.separators))
	copy(separatorsCopy, c.separators)

	return decoratorsCopy, multiplexersCopy, separatorsCopy
}

func (c *conveyer) runDecoratorHandlers(group *errgroup.Group, ctx context.Context, decorators []decoratorSpec) {
	for _, decorator := range decorators {
		dec := decorator

		group.Go(func() error {
			input, _ := c.getChannel(dec.input)
			output, _ := c.getChannel(dec.output)

			return dec.fn(ctx, input, output)
		})
	}
}

func (c *conveyer) runMultiplexerHandlers(group *errgroup.Group, ctx context.Context, multiplexers []multiplexerSpec) {
	for _, multiplexer := range multiplexers {
		mux := multiplexer

		group.Go(func() error {
			inputs := make([]chan string, len(mux.inputs))

			for i, name := range mux.inputs {
				inputs[i], _ = c.getChannel(name)
			}

			output, _ := c.getChannel(mux.output)

			return mux.fn(ctx, inputs, output)
		})
	}
}

func (c *conveyer) runSeparatorHandlers(group *errgroup.Group, ctx context.Context, separators []separatorSpec) {
	for _, separator := range separators {
		sep := separator

		group.Go(func() error {
			input, _ := c.getChannel(sep.input)

			outputs := make([]chan string, len(sep.outputs))

			for i, name := range sep.outputs {
				outputs[i], _ = c.getChannel(name)
			}

			return sep.fn(ctx, input, outputs)
		})
	}
}

func (c *conveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case channel <- data:
		return nil
	case <-time.After(timeoutTime):
		return ErrTimeout
	default:
		return ErrFullChannel
	}
}

func (c *conveyer) Recv(output string) (string, error) {
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
