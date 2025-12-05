package conveyor

import (
	"context"
)

type Сonveyor struct {
	channels map[string]chan string
	handlers []func(ctx context.Context) error
	size     int
}

func New(size int) Сonveyor {
	return Сonveyor{
		channels: make(map[string]chan string),
		handlers: make([]func(ctx context.Context) error, 0),
		size:     size,
	}
}

func (c *Сonveyor) createChan(chName string) chan string {
	if channel, exists := c.channels[chName]; exists {
		return channel
	}

	newChannel := make(chan string, c.size)
	c.channels[chName] = newChannel

	return newChannel
}

func (c *Сonveyor) RegisterDecorator(
	fn func(cntx context.Context, in chan string, out chan string) error,
	in string,
	out string) {

}

func (c *Сonveyor) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
)
func (c *Сonveyor) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
)

func Run(ctx context.Context) error {
	return nil
}

func Send(input string, data string) error {
	return nil
}

func Recv(output string) (string, error) {
	return "nil", nil
}
