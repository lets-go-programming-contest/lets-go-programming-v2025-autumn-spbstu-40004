package conveyer

import (
	"context"
	"sync"
)

type conveyer struct {
	mu       sync.RWMutex
	channels map[string]chan string
	handlers []handlerFunc
	size     int
}

type handlerFunc struct {
	fn      interface{}
	inputs  []string
	outputs []string
}

func New(size int) *conveyer {
	return &conveyer{
		channels: make(map[string]chan string),
		handlers: make([]handlerFunc, 0),
		size:     size,
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
