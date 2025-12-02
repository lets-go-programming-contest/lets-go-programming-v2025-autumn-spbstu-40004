package conveyer

import (
	"context"
)

type Conveyer struct {
	chans        map[string]chan string
	size         int
	decorators   []decorator
	multiplexers []multiplexer
	separators   []separator
}

type decorator struct {
	fn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error
	input  string
	output string
}

type multiplexer struct {
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error
	inputs []string
	output string
}

type separator struct {
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error
	input   string
	outputs []string
}

func New(size int) *Conveyer {
	return &Conveyer{
		chans:        make(map[string]chan string),
		size:         size,
		decorators:   make([]decorator, 0),
		multiplexers: make([]multiplexer, 0),
		separators:   make([]separator, 0),
	}
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
	if _, ok := c.chans[input]; !ok {
		channel := make(chan string, c.size)
		c.chans[input] = channel
	}

	if _, ok := c.chans[output]; !ok {
		channel := make(chan string, c.size)
		c.chans[output] = channel
	}

	c.decorators = append(c.decorators, decorator{
		fn:     fn,
		input:  input,
		output: output,
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
	for _, input := range inputs {
		if _, ok := c.chans[input]; !ok {
			channel := make(chan string, c.size)
			c.chans[input] = channel
		}
	}

	if _, ok := c.chans[output]; !ok {
		channel := make(chan string, c.size)
		c.chans[output] = channel
	}

	c.multiplexers = append(c.multiplexers, multiplexer{
		fn:     fn,
		inputs: inputs,
		output: output,
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
	if _, ok := c.chans[input]; !ok {
		channel := make(chan string, c.size)
		c.chans[input] = channel
	}

	for _, output := range outputs {
		if _, ok := c.chans[output]; !ok {
			channel := make(chan string, c.size)
			c.chans[output] = channel
		}
	}

	c.separators = append(c.separators, separator{
		fn:      fn,
		input:   input,
		outputs: outputs,
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	return nil
}

func (c *Conveyer) Send(input string, data string) error {
	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	return "", nil
}
