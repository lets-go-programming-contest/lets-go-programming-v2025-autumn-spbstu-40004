package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrChanNotFound = errors.New("chan not found")

const ClosedChan string = "undefined"

type Conveyer struct {
	chans    map[string]chan string
	chanSize int
	handlers []func(ctx context.Context) error
	rwMtx    sync.RWMutex
}

func New(size int) Conveyer {
	return Conveyer{
		chans:    make(map[string]chan string, 0),
		chanSize: size,
		handlers: make([]func(ctx context.Context) error, 0),
		rwMtx:    sync.RWMutex{},
	}
}

func (c *Conveyer) createChan(chName string) chan string {
	if ch, exists := c.chans[chName]; exists {
		return ch
	}

	ch := make(chan string, c.chanSize)
	c.chans[chName] = ch

	return ch
}

func (c *Conveyer) RegisterDecorator(
	handler func(
		ctx context.Context,
		input, output chan string,
	) error,
	input, output string,
) {
	c.rwMtx.Lock()
	defer c.rwMtx.Unlock()

	ichan := c.createChan(input)
	ochan := c.createChan(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handler(ctx, ichan, ochan)
	})
}

func (c *Conveyer) RegisterMultiplexer(
	handler func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	c.rwMtx.Lock()
	defer c.rwMtx.Unlock()

	ichans := make([]chan string, 0)
	for _, input := range inputs {
		ichans = append(ichans, c.createChan(input))
	}

	ochan := c.createChan(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handler(ctx, ichans, ochan)
	})
}

func (c *Conveyer) RegisterSeparator(
	handler func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	c.rwMtx.Lock()
	defer c.rwMtx.Unlock()

	ichan := c.createChan(input)

	ochans := make([]chan string, 0)
	for _, output := range outputs {
		ochans = append(ochans, c.createChan(output))
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handler(ctx, ichan, ochans)
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer func() {
		c.rwMtx.RLock()
		defer c.rwMtx.RUnlock()

		for _, ch := range c.chans {
			close(ch)
		}
	}()

	c.rwMtx.RLock()

	errGr, ctx := errgroup.WithContext(ctx)
	for _, fn := range c.handlers {
		errGr.Go(func() error {
			return fn(ctx)
		})
	}

	c.rwMtx.RUnlock()

	if err := errGr.Wait(); err != nil {
		return fmt.Errorf("handler error received: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	c.rwMtx.RLock()
	ichan, exists := c.chans[input]
	c.rwMtx.RUnlock()

	if !exists {
		return ErrChanNotFound
	}

	ichan <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	c.rwMtx.RLock()
	ochan, exists := c.chans[output]
	c.rwMtx.RUnlock()

	if !exists {
		return "", ErrChanNotFound
	}

	data, ok := <-ochan
	if !ok {
		return ClosedChan, nil
	}

	return data, nil
}
