package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var ErrChanNotFound = errors.New("chan not found")

type conveyer struct {
	mu          sync.RWMutex
	channels    map[string]chan string
	handlers    []handlerFunc
	size        int
	isRunning   bool
	cancelFuncs []context.CancelFunc
}

type handlerFunc struct {
	fn      interface{}
	inputs  []string
	outputs []string
}

func New(size int) *conveyer {
	return &conveyer{
		channels:    make(map[string]chan string),
		handlers:    make([]handlerFunc, 0),
		size:        size,
		cancelFuncs: make([]context.CancelFunc, 0),
	}
}

func (c *conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

func (c *conveyer) getChannel(name string) (chan string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch, exists := c.channels[name]
	return ch, exists
}

func (c *conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerFunc{
		fn:      fn,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, input := range inputs {
		c.getOrCreateChannel(input)
	}
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerFunc{
		fn:      fn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.getOrCreateChannel(input)
	for _, output := range outputs {
		c.getOrCreateChannel(output)
	}

	c.handlers = append(c.handlers, handlerFunc{
		fn:      fn,
		inputs:  []string{input},
		outputs: outputs,
	})
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.isRunning {
		c.mu.Unlock()
		return fmt.Errorf("conveyer is already running")
	}
	c.isRunning = true
	c.mu.Unlock()

	defer c.cleanup()

	handlerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	errorCh := make(chan error, len(c.handlers))

	for _, handler := range c.handlers {
		wg.Add(1)
		go func(h handlerFunc) {
			defer wg.Done()

			var err error
			switch fn := h.fn.(type) {
			case func(ctx context.Context, input chan string, output chan string) error:
				inputChan := c.getOrCreateChannel(h.inputs[0])
				outputChan := c.getOrCreateChannel(h.outputs[0])
				err = fn(handlerCtx, inputChan, outputChan)
			case func(ctx context.Context, inputs []chan string, output chan string) error:
				inputChans := make([]chan string, len(h.inputs))
				for i, input := range h.inputs {
					inputChans[i] = c.getOrCreateChannel(input)
				}
				outputChan := c.getOrCreateChannel(h.outputs[0])
				err = fn(handlerCtx, inputChans, outputChan)
			case func(ctx context.Context, input chan string, outputs []chan string) error:
				inputChan := c.getOrCreateChannel(h.inputs[0])
				outputChans := make([]chan string, len(h.outputs))
				for i, output := range h.outputs {
					outputChans[i] = c.getOrCreateChannel(output)
				}
				err = fn(handlerCtx, inputChan, outputChans)
			default:
				err = fmt.Errorf("unknown handler type")
			}

			if err != nil {
				select {
				case errorCh <- err:
				default:
				}
				cancel()
			}
		}(handler)
	}

	go func() {
		wg.Wait()
		close(errorCh)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err, ok := <-errorCh:
		if ok && err != nil {
			return err
		}
	}

	return nil
}

func (c *conveyer) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}
	c.handlers = make([]handlerFunc, 0)
	c.isRunning = false
	c.cancelFuncs = make([]context.CancelFunc, 0)
}

func (c *conveyer) Send(input string, data string) error {
	ch, exists := c.getChannel(input)
	if !exists {
		return ErrChanNotFound
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	ch, exists := c.getChannel(output)
	if !exists {
		return "", ErrChanNotFound
	}

	select {
	case data, ok := <-ch:
		if !ok {
			return "undefined", nil
		}
		return data, nil
	default:
		return "", errors.New("no data available")
	}
}
