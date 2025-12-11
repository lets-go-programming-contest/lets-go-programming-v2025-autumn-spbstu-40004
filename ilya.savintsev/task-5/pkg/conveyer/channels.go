package conveyer

func (c *DefaultConveyer) obtainChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, ok := c.channels[name]
	if !ok {
		if c.bufferSize <= 0 {
			ch = make(chan string)
		} else {
			ch = make(chan string, c.bufferSize)
		}
		c.channels[name] = ch
	}
	return ch
}

func (c *DefaultConveyer) getChannel(name string) (chan string, error) {
	c.mu.RLock()
	ch, ok := c.channels[name]
	c.mu.RUnlock()

	if !ok {
		return nil, ErrChanNotFound
	}
	return ch, nil
}
