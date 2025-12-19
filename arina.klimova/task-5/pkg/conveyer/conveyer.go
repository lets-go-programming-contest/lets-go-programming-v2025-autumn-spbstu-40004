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
	ErrChanNotFound     = errors.New("chan not found")
	ErrTimeout          = errors.New("timeout")
	ErrFullChannel      = errors.New("channel is full")
	ErrAlreadyRunning   = errors.New("conveyer is already running")
	ErrCannotRegister   = errors.New("cannot register handler while conveyer is running")
	ErrChannelClosed    = errors.New("channel closed")
	ErrConveyerFinished = errors.New("conveyer finished")
)

const (
	timeoutTime = 100 * time.Millisecond
)

type conveyer struct {
	mu            sync.RWMutex
	channels      map[string]chan string
	bufferSize    int
	decorators    []decoratorSpec
	multiplexers  []multiplexerSpec
	separators    []separatorSpec
	runInProgress bool
	runStarted    bool
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
		channels:     make(map[string]chan string),
		bufferSize:   size,
		decorators:   make([]decoratorSpec, 0),
		multiplexers: make([]multiplexerSpec, 0),
		separators:   make([]separatorSpec, 0),
	}
}

func (c *conveyer) obtainChannel(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

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
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.runStarted {
		return ErrCannotRegister
	}

	if _, exists := c.channels[input]; !exists {
		c.channels[input] = make(chan string, c.bufferSize)
	}
	if _, exists := c.channels[output]; !exists {
		c.channels[output] = make(chan string, c.bufferSize)
	}

	c.decorators = append(c.decorators, decoratorSpec{
		fn:     decoratorFunc,
		input:  input,
		output: output,
	})

	return nil
}

func (c *conveyer) RegisterMultiplexer(
	multiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.runStarted {
		return ErrCannotRegister
	}

	for _, inputName := range inputs {
		if _, exists := c.channels[inputName]; !exists {
			c.channels[inputName] = make(chan string, c.bufferSize)
		}
	}

	if _, exists := c.channels[output]; !exists {
		c.channels[output] = make(chan string, c.bufferSize)
	}

	c.multiplexers = append(c.multiplexers, multiplexerSpec{
		fn:     multiplexerFunc,
		inputs: inputs,
		output: output,
	})

	return nil
}

func (c *conveyer) RegisterSeparator(
	separatorFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.runInProgress {
		return ErrCannotRegister
	}

	if _, exists := c.channels[input]; !exists {
		c.channels[input] = make(chan string, c.bufferSize)
	}

	for _, outputName := range outputs {
		c.obtainChannel(outputName)
	}

	c.separators = append(c.separators, separatorSpec{
		fn:      separatorFunc,
		input:   input,
		outputs: outputs,
	})

	return nil
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.runInProgress {
		c.mu.Unlock()
		return ErrAlreadyRunning
	}

	decoratorsCopy := make([]decoratorSpec, len(c.decorators))
	copy(decoratorsCopy, c.decorators)

	multiplexersCopy := make([]multiplexerSpec, len(c.multiplexers))
	copy(multiplexersCopy, c.multiplexers)

	separatorsCopy := make([]separatorSpec, len(c.separators))
	copy(separatorsCopy, c.separators)

	c.runInProgress = true
	c.runStarted = true
	c.mu.Unlock()

	group, groupCtx := errgroup.WithContext(ctx)

	for _, decorator := range decoratorsCopy {
		dec := decorator

		group.Go(func() error {
			input, err := c.getChannel(dec.input)
			if err != nil {
				return err
			}

			output, err := c.getChannel(dec.output)
			if err != nil {
				return err
			}

			return dec.fn(groupCtx, input, output)
		})
	}

	for _, multiplexer := range multiplexersCopy {
		mux := multiplexer

		group.Go(func() error {
			inputs := make([]chan string, len(mux.inputs))
			for i, name := range mux.inputs {
				channel, err := c.getChannel(name)
				if err != nil {
					return err
				}
				inputs[i] = channel
			}

			output, err := c.getChannel(mux.output)
			if err != nil {
				return err
			}

			return mux.fn(groupCtx, inputs, output)
		})
	}

	for _, separator := range separatorsCopy {
		sep := separator

		group.Go(func() error {
			input, err := c.getChannel(sep.input)
			if err != nil {
				return err
			}

			outputs := make([]chan string, len(sep.outputs))
			for i, name := range sep.outputs {
				channel, err := c.getChannel(name)
				if err != nil {
					return err
				}
				outputs[i] = channel
			}

			return sep.fn(groupCtx, input, outputs)
		})
	}

	err := group.Wait()

	c.mu.Lock()
	for _, ch := range c.channels {
		close(ch)
	}
	c.runInProgress = false
	c.mu.Unlock()

	if err != nil {
		return fmt.Errorf("%w: %v", ErrConveyerFinished, err)
	}

	return nil
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

	select {
	case data, ok := <-channel:
		if !ok {
			return "undefined", nil
		}
		return data, nil
	case <-time.After(timeoutTime):
		return "", ErrTimeout
	}
}
