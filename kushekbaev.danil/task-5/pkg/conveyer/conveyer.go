package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

const undefined = "undefined"

var ErrChannelNotFound = errors.New("channel not found")

type Conveyer struct {
	size     int
	channels map[string]chan string
	handlers []func(context.Context) error
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]func(context.Context) error, 0),
	}
}

func (cnv *Conveyer) makeChannel(name string) {
	if _, ok := cnv.channels[name]; !ok {
		cnv.channels[name] = make(chan string, cnv.size)
	}
}

func (cnv *Conveyer) makeChannels(names ...string) {
	for _, n := range names {
		cnv.makeChannel(n)
	}
}
