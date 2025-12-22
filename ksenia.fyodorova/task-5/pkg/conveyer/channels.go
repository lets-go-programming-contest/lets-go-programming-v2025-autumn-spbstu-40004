package conveyer

import (
	"errors"
	"sync"
)

var ErrChanNotFound = errors.New("chan not found")

type DefaultConveyer struct {
	mu           sync.RWMutex
	channels     map[string]chan string
	bufferSize   int
	decorators   []specDecorator
	multiplexers []specMultiplexer
	separators   []specSeparator
}

func (c *DefaultConveyer) obtainChannel(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.channels[name]; exists {
		return
	}

	channel := make(chan string, c.bufferSize)
	c.channels[name] = channel
}

func (c *DefaultConveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if channel, exists := c.channels[name]; exists {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *DefaultConveyer) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		close(channel)
	}
}
