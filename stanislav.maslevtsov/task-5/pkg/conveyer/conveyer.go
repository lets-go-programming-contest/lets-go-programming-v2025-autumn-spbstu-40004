package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var ErrChanNotFound = errors.New("chan not found")

const ClosedChan string = "undefined"

type Conveyer struct {
	chans    map[string]chan string
	chanSize int
	handlers []func(ctx context.Context) error
}

func New(size int) Conveyer {
	return Conveyer{
		chans:    make(map[string]chan string, 0),
		chanSize: size,
		handlers: make([]func(ctx context.Context) error, 0),
	}
}

func (c *Conveyer) createChan(name string) chan string {
	if ch, ok := c.chans[name]; ok {
		return ch
	}

	ch := make(chan string, c.chanSize)
	c.chans[name] = ch

	return ch
}

func (c *Conveyer) RegisterDecorator(
	fn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, c.createChan(input), c.createChan(output))
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	var ichans []chan string
	for _, input := range inputs {
		ichans = append(ichans, c.createChan(input))
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, ichans, c.createChan(output))
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	var ochans []chan string
	for _, output := range outputs {
		ochans = append(ochans, c.createChan(output))
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, c.createChan(input), ochans)
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer func() {
		for _, ch := range c.chans {
			close(ch)
		}
	}()

	errGr, ctx := errgroup.WithContext(ctx)
	for _, fn := range c.handlers {
		errGr.Go(func() error {
			return fn(ctx)
		})
	}

	if err := errGr.Wait(); err != nil {
		return fmt.Errorf("handler error received: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	ch, ok := c.chans[input]
	if !ok {
		return ErrChanNotFound
	}

	ch <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	ch, ok := c.chans[output]
	if !ok {
		return "", ErrChanNotFound
	}

	data, ok := <-ch
	if !ok {
		return ClosedChan, nil
	}

	return data, nil
}
