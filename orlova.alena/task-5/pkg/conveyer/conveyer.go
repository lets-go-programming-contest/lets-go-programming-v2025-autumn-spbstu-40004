package conveyer

import (
	"fmt"
)

type Conveyer struct {
	channels map[string]chan string
	stop     chan struct{}
}

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		stop:     make(chan struct{}),
	}
}

func (conv *Conveyer) createChannel(id string) {
	if _, exists := conv.channels[id]; !exists {
		conv.channels[id] = make(chan string, 10)
	}
}

func (conv *Conveyer) getChannel(id string) (chan string, error) {
	ch, exists := conv.channels[id]
	if !exists {
		return nil, fmt.Errorf("chan not found")
	}
	return ch, nil
}
