package conveyer

import (
	"fmt"
	"sync"
)

type ChanManager struct {
	mu    sync.RWMutex
	chans map[string]chan string
}

func NewChanManager() *ChanManager {
	return &ChanManager{
		chans: make(map[string]chan string),
	}
}

func (cm *ChanManager) GetOrCreate(name string, size int) chan string {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if ch, ok := cm.chans[name]; ok {
		return ch
	}

	ch := make(chan string, size)
	cm.chans[name] = ch
	return ch
}

func (cm *ChanManager) Get(name string) (chan string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	ch, ok := cm.chans[name]
	return ch, ok
}

func (cm *ChanManager) CloseAll() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for name, ch := range cm.chans {
		close(ch)
		delete(cm.chans, name)
	}
}

func (cm *ChanManager) Send(name string, data string) error {
	cm.mu.RLock()
	ch, ok := cm.chans[name]
	cm.mu.RUnlock()

	if !ok {
		return fmt.Errorf("chan not found")
	}

	select {
	case ch <- data:
		return nil
	default:
		return fmt.Errorf("chan is full")
	}
}

func (cm *ChanManager) Recv(name string) (string, error) {
	cm.mu.RLock()
	ch, ok := cm.chans[name]
	cm.mu.RUnlock()

	if !ok {
		return "", fmt.Errorf("chan not found")
	}

	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return data, nil
}
