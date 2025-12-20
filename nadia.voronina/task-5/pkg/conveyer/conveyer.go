package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrNoChannels = errors.New("chan not found")

type IConveyer interface {
	RegisterDecorator(
		fn func(
			ctx context.Context,
			input chan string,
			output chan string,
		) error,
		input string,
		output string,
	)
	RegisterMultiplexer(
		fn func(
			ctx context.Context,
			inputs []chan string,
			output chan string,
		) error,
		inputs []string,
		output string,
	)
	RegisterSeparator(
		fn func(
			ctx context.Context,
			input chan string,
			outputs []chan string,
		) error,
		input string,
		outputs []string,
	)
	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

type Conveyer struct {
	size     int
	channels map[string]chan string
	workers  []func(ctx context.Context) error
	rwlock   sync.RWMutex
}

func (c *Conveyer) RegisterDecorator(
	decoratorFunction func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()
	inCh := c.getOrCreateChannel(input)
	outCh := c.getOrCreateChannel(output)
	worker := func(ctx context.Context) error {
		return decoratorFunction(ctx, inCh, outCh)
	}

	c.workers = append(c.workers, worker)
}

func (c *Conveyer) RegisterMultiplexer(
	multiplexerFunction func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()

	inChans := make([]chan string, len(inputs))
	for i, name := range inputs {
		inChans[i] = c.getOrCreateChannel(name)
	}

	outCh := c.getOrCreateChannel(output)
	worker := func(ctx context.Context) error {
		return multiplexerFunction(ctx, inChans, outCh)
	}

	c.workers = append(c.workers, worker)
}

func (c *Conveyer) RegisterSeparator(
	separatorFunction func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()
	inCh := c.getOrCreateChannel(input)
	outChans := make([]chan string, len(outputs))

	for i, name := range outputs {
		outChans[i] = c.getOrCreateChannel(name)
	}

	worker := func(ctx context.Context) error {
		return separatorFunction(ctx, inCh, outChans)
	}

	c.workers = append(c.workers, worker)
}

func (c *Conveyer) Run(ctx context.Context) error {
	if len(c.workers) == 0 {
		return nil
	}

	defer func() {
		c.rwlock.Lock()
		defer c.rwlock.Unlock()

		for _, ch := range c.channels {
			close(ch)
		}
	}()

	errorgroup, ctx := errgroup.WithContext(ctx)

	for _, worker := range c.workers {
		w := worker

		errorgroup.Go(func() error {
			return w(ctx)
		})
	}

	if err := errorgroup.Wait(); err != nil {
		return fmt.Errorf("handler error received: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()

	ch, exists := c.channels[input]
	if !exists {
		return ErrNoChannels
	}
	ch <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()

	ch, exists := c.channels[output]
	if !exists {
		return "", ErrNoChannels
	}

	val, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return val, nil
}

func (c *Conveyer) getOrCreateChannel(name string) chan string {
	if c.channels == nil {
		c.channels = make(map[string]chan string)
	}

	workerChannel, exists := c.channels[name]
	if !exists {
		workerChannel = make(chan string, c.size)
		c.channels[name] = workerChannel
	}

	return workerChannel
}

func New(size int) Conveyer {
	return Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		workers:  []func(ctx context.Context) error{},
		rwlock:   sync.RWMutex{},
	}
}
