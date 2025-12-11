package conveyer

func (c *DefaultConveyer) obtainChannel(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.channels[name]
	if exists {
		return
	}

	var channel chan string
	if c.bufferSize > 0 {
		channel = make(chan string, c.bufferSize)
	} else {
		channel = make(chan string)
	}

	c.channels[name] = channel
}

func (c *DefaultConveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	channel, exists := c.channels[name]
	c.mu.RUnlock()

	if !exists {
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
