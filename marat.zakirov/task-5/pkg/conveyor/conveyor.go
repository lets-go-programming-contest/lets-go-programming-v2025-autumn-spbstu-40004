package conveyor

import (
	"context"
)

type Conveyor struct {
	channels map[string]chan string
	handlers []func(ctx context.Context) error
	size     int
}

func New(size int) Conveyor {
	return Conveyor{
		channels: make(map[string]chan string),
		handlers: make([]func(ctx context.Context) error, 0),
		size:     size,
	}
}

func (c *Conveyor) createChannel(chName string) chan string {
	if channel, exists := c.channels[chName]; exists {
		return channel
	}

	newChannel := make(chan string, c.size)
	c.channels[chName] = newChannel

	return newChannel
}

func (c *Conveyor) RegisterDecorator(
	fn func(cntx context.Context, in chan string, out chan string) error,
	in string,
	out string,
) {
	inChannel := c.createChannel(in)
	outChannel := c.createChannel(out)

	c.handlers = append(c.handlers, func(cntx context.Context) error {
		return fn(cntx, inChannel, outChannel)
	})
}

func (c *Conveyor) RegisterMultiplexer(
	fn func(ctx context.Context, ins []chan string, out chan string) error,
	ins []string,
	out string,
) {
	inChannels := make([]chan string, len(ins))
	for i, name := range ins {
		inChannels[i] = c.createChannel(name)
	}
	outChannel := c.createChannel(out)

	c.handlers = append(c.handlers, func(cntx context.Context) error {
		return fn(cntx, inChannels, outChannel)
	})
}

func (c *Conveyor) RegisterSeparator(
	fn func(ctx context.Context, in chan string, outs []chan string) error,
	in string,
	outs []string,
) {
	inChannel := c.createChannel(in)
	outChannels := make([]chan string, len(outs))
	for i, name := range outs {
		outChannels[i] = c.createChannel(name)
	}

	c.handlers = append(c.handlers, func(cntx context.Context) error {
		return fn(cntx, inChannel, outChannels)
	})
}

func Run(ctx context.Context) error {
	return nil
}

func Send(input string, data string) error {
	return nil
}

func Recv(output string) (string, error) {
	return "nil", nil
}
