package conveyer

func (c *DefaultConveyer) obtainChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	channel, exists := c.channels[name]
	if !exists {
		if c.bufferSize <= 0 {
			channel = make(chan string)
		} else {
			channel = make(chan string, c.bufferSize)
		}
		c.channels[name] = channel
	}

	return channel
}

func (c *DefaultConveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	channel, ok := c.channels[name]
	c.mu.RUnlock()

	if !ok {
		return nil, ErrChanNotFound
	}

	return channel, nil
}

func (c *DefaultConveyer) closeAllChannels() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		for _, channel := range c.channels {
			if channel != nil {
				close(channel)
			}
		}
	})
}
