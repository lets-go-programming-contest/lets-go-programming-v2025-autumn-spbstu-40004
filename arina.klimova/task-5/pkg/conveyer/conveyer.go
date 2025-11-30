package conveyer

import (
	"context"
	"sync"
)

type conveyer interface {
	RegisterDecorator(
		fn func(ctx context.Context, input chan string, output chan string) error,
		input string,
		output string,
	)
	RegisterMultiplexer(
		fn func(ctx context.Context, inputs []chan string, output chan string) error,
		inputs []string,
		output string,
	)
	RegisterSeparator(
		fn func(ctx context.Context, input chan string, outputs []chan string) error,
		input string,
		outputs []string,
	)
	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

type conveyerImpl struct {
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

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		channels: make(map[string]chan string),
		handlers: make([]handlerFunc, 0),
		size:     size,
	}
}
