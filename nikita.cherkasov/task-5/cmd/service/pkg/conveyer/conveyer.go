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
