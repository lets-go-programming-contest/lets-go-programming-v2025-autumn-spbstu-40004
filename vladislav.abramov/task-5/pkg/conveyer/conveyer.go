package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound        = errors.New("chan not found")
	ErrAlreadyRunning      = errors.New("conveyer is already running")
	ErrChannelFullOrClosed = errors.New("channel is full or closed")
	ErrNoDataAvailable     = errors.New("no data available")
)

const errUndefined = "undefined"

type conveyer struct {
	mu           sync.RWMutex
	channels     map[string]chan string
	size         int
	decorators   []decoratorConfig
	multiplexers []multiplexerConfig
	separators   []separatorConfig
	isRunning    bool
	runCancel    context.CancelFunc
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
		mu:           sync.RWMutex{},
		channels:     make(map[string]chan string),
		size:         size,
		decorators:   make([]decoratorConfig, 0),
		multiplexers: make([]multiplexerConfig, 0),
		separators:   make([]separatorConfig, 0),
		isRunning:    false,
		runCancel:    nil,
	}
}

func (c *conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.size)
	c.channels[name] = channel

	return channel
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

	if _, exists := c.channels[input]; !exists {
		c.channels[input] = make(chan string, c.size)
	}

	if _, exists := c.channels[output]; !exists {
		c.channels[output] = make(chan string, c.size)
	}

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
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, inputName := range inputs {
		if _, exists := c.channels[inputName]; !exists {
			c.channels[inputName] = make(chan string, c.size)
		}
	}

	if _, exists := c.channels[output]; !exists {
		c.channels[output] = make(chan string, c.size)
	}

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
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.channels[input]; !exists {
		c.channels[input] = make(chan string, c.size)
	}

	for _, outputName := range outputs {
		if _, exists := c.channels[outputName]; !exists {
			c.channels[outputName] = make(chan string, c.size)
		}
	}

	c.separators = append(c.separators, separatorConfig{
		fn:      separatorFunc,
		input:   input,
		outputs: outputs,
	})
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.isRunning {
		c.mu.Unlock()
		return ErrAlreadyRunning
	}

	runCtx, cancel := context.WithCancel(ctx)
	c.runCancel = cancel
	c.isRunning = true
	c.mu.Unlock()

	defer func() {
		c.closeAllChannels()

		c.mu.Lock()
		c.isRunning = false
		c.runCancel = nil
		c.mu.Unlock()
	}()

	errorGroup, groupCtx := errgroup.WithContext(runCtx)

	c.runDecorators(groupCtx, errorGroup)
	c.runMultiplexers(groupCtx, errorGroup)
	c.runSeparators(groupCtx, errorGroup)

	if err := errorGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *conveyer) runDecorators(ctx context.Context, errorGroup *errgroup.Group) {
	c.mu.RLock()
	decorators := make([]decoratorConfig, len(c.decorators))
	copy(decorators, c.decorators)
	c.mu.RUnlock()

	for _, decorator := range decorators {
		inputChannel, _ := c.getChannel(decorator.input)
		outputChannel, _ := c.getChannel(decorator.output)

		currentDecorator := decorator

		errorGroup.Go(func() error {
			return currentDecorator.fn(ctx, inputChannel, outputChannel)
		})
	}
}

func (c *conveyer) runMultiplexers(ctx context.Context, errorGroup *errgroup.Group) {
	c.mu.RLock()
	multiplexers := make([]multiplexerConfig, len(c.multiplexers))
	copy(multiplexers, c.multiplexers)
	c.mu.RUnlock()

	for _, multiplexer := range multiplexers {
		inputChannels := make([]chan string, len(multiplexer.inputs))

		for index, inputName := range multiplexer.inputs {
			inputChannels[index], _ = c.getChannel(inputName)
		}

		outputChannel, _ := c.getChannel(multiplexer.output)

		currentMultiplexer := multiplexer

		errorGroup.Go(func() error {
			return currentMultiplexer.fn(ctx, inputChannels, outputChannel)
		})
	}
}

func (c *conveyer) runSeparators(ctx context.Context, errorGroup *errgroup.Group) {
	c.mu.RLock()
	separators := make([]separatorConfig, len(c.separators))
	copy(separators, c.separators)
	c.mu.RUnlock()

	for _, separator := range separators {
		inputChannel, _ := c.getChannel(separator.input)
		outputChannels := make([]chan string, len(separator.outputs))

		for index, outputName := range separator.outputs {
			outputChannels[index], _ = c.getChannel(outputName)
		}

		currentSeparator := separator

		errorGroup.Go(func() error {
			return currentSeparator.fn(ctx, inputChannel, outputChannels)
		})
	}
}

func (c *conveyer) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		close(channel)
	}

	c.channels = make(map[string]chan string)
}

func (c *conveyer) Send(input string, data string) error {
	c.mu.RLock()
	if !c.isRunning {
		c.mu.RUnlock()
		return errors.New("conveyer is not running")
	}
	c.mu.RUnlock()

	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case channel <- data:
		return nil
	default:
		return ErrChannelFullOrClosed
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	c.mu.RLock()
	if !c.isRunning {
		c.mu.RUnlock()
		return "", errors.New("conveyer is not running")
	}
	c.mu.RUnlock()

	channel, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	select {
	case data, ok := <-channel:
		if !ok {
			return errUndefined, nil
		}
		return data, nil
	default:
		return "", ErrNoDataAvailable
	}
}
