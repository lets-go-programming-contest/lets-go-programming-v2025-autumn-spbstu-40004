package conveyer

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
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
}

func (c *Conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
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
	workers := make([]func(context.Context) error, len(c.workers))
	copy(workers, c.workers)

	if len(workers) == 0 {
		return nil
	}

	defer func() {
		for _, ch := range c.channels {
			close(ch)
		}
	}()

	var eg errgroup.Group

	for _, worker := range workers {
		w := worker
		eg.Go(func() error {
			return w(ctx)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	ch, exists := c.channels[input]
	if !exists {
		return fmt.Errorf("chan not found")
	}
	ch <- data
	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	ch, exists := c.channels[output]
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
	}
}
