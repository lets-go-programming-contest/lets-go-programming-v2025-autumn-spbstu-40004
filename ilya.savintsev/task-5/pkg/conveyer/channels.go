package conveyer

func (c *DefaultConveyer) obtainChannel(name string) {
	if _, exists := c.channels[name]; exists {
		return
	}

	channel := make(chan string, c.bufferSize)
	c.channels[name] = channel
}

func (c *DefaultConveyer) getChannel(name string) (chan string, error) {
	if channel, exists := c.channels[name]; exists {
		return channel, nil
	}

	return nil, ErrChanNotFound
}

func (c *DefaultConveyer) closeAllChannels() {
	for _, channel := range c.channels {
		close(channel)
	}
}
