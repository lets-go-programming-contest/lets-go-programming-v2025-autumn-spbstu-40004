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

func Run(ctx context.Context) error {
	return nil
}

func Send(input string, data string) error {
	return nil
}

func Recv(output string) (string, error) {
	return "nil", nil
}
