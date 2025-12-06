package conveyer

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrChanNotFound = errors.New("chan not found")
	ErrChanFull     = errors.New("channel is full")
)

const undefinedData = "undefined"

type conveyerImpl struct {
	size     int
	channels map[string]chan string
	handlers []func(context.Context) error
	mu       sync.RWMutex
}

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]func(context.Context) error, 0),
		mu:       sync.RWMutex{},
	}
}

func (conveyer *conveyerImpl) getOrCreateChannel(name string) chan string {
	if existingChannel, exists := conveyer.channels[name]; exists {
		return existingChannel
	}

	newChannel := make(chan string, conveyer.size)
	conveyer.channels[name] = newChannel

	return newChannel
}

func (conveyer *conveyerImpl) getOrCreateChannels(names ...string) {
	for _, name := range names {
		conveyer.getOrCreateChannel(name)
	}
}

func (conveyer *conveyerImpl) RegisterDecorator(
	decoratorFunc func(ctx context.Context, input chan string, output chan string) error,
	inputName string,
	outputName string,
) {
	conveyer.mu.Lock()
	defer conveyer.mu.Unlock()

	conveyer.getOrCreateChannels(inputName, outputName)

	inputChannel := conveyer.channels[inputName]
	outputChannel := conveyer.channels[outputName]

	conveyer.handlers = append(conveyer.handlers, func(ctx context.Context) error {
		return decoratorFunc(ctx, inputChannel, outputChannel)
	})
}
