package conveyer

import (
	"context"
	"fmt"
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
	if len(c.workers) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, len(c.workers))
	done := make(chan struct{})

	for _, worker := range c.workers {
		go func(w func(context.Context) error) {
			err := w(ctx)
			if err != nil {
				errCh <- err
			} else {
				errCh <- nil
			}
		}(worker)
	}

	go func() {
		for i := 0; i < len(c.workers); i++ {
			if err := <-errCh; err != nil {
				cancel()
				break
			}
		}
		close(done)
	}()

	select {
	case <-ctx.Done():
	case <-done:
	}

	for _, ch := range c.channels {
		close(ch)
	}

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	default:
	}

	return ctx.Err()
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
		c.channels = make(map[string]chan string)
	}
	ch, exists := c.channels[name]
	if !exists {
		ch = make(chan string)
		c.channels[name] = ch
	}
	return ch
}
