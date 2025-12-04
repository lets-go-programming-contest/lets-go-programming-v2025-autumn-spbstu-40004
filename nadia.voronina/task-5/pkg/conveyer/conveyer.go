package conveyer

import (
	"context"
	"fmt"
	"sync"
)

type conveyer interface {
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
	mu       sync.RWMutex
}

func (c *Conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	inCh := c.getOrCreateChannel(input)
	outCh := c.getOrCreateChannel(output)
	worker := func(ctx context.Context) error {
		return fn(ctx, inCh, outCh)
	}
	c.workers = append(c.workers, worker)
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	inChans := make([]chan string, len(inputs))
	for i, name := range inputs {
		inChans[i] = c.getOrCreateChannel(name)
	}
	outCh := c.getOrCreateChannel(output)
	worker := func(ctx context.Context) error {
		return fn(ctx, inChans, outCh)
	}
	c.workers = append(c.workers, worker)
}

func (c *Conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	inCh := c.getOrCreateChannel(input)
	outChans := make([]chan string, len(outputs))
	for i, name := range outputs {
		outChans[i] = c.getOrCreateChannel(name)
	}
	worker := func(ctx context.Context) error {
		return fn(ctx, inCh, outChans)
	}
	c.workers = append(c.workers, worker)
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.mu.RLock()
	workers := make([]func(context.Context) error, len(c.workers))
	copy(workers, c.workers)
	c.mu.RUnlock()

	if len(workers) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, len(workers))

	var wg sync.WaitGroup
	wg.Add(len(workers))

	for _, worker := range workers {
		go func(w func(context.Context) error) {
			defer wg.Done()
			err := w(ctx)
			errCh <- err
			cancel()
		}(worker)
	}

	wg.Wait()

	c.mu.Lock()
	for _, ch := range c.channels {
		close(ch)
	}
	c.mu.Unlock()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	default:
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	c.mu.RLock()
	ch, exists := c.channels[input]
	c.mu.RUnlock()
	if !exists {
		return fmt.Errorf("chan not found")
	}
	ch <- data
	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	c.mu.RLock()
	ch, exists := c.channels[output]
	c.mu.RUnlock()
	if !exists {
		return "", fmt.Errorf("chan not found")
	}
	val, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return val, nil
}

func (c *Conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.channels == nil {
		c.channels = make(map[string]chan string, c.size)
	}
	ch, exists := c.channels[name]
	if !exists {
		ch = make(chan string)
		c.channels[name] = ch
	}
	return ch
}

func New(size int) Conveyer {
	return Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		workers:  []func(ctx context.Context) error{},
		mu:       sync.RWMutex{},
	}
}
