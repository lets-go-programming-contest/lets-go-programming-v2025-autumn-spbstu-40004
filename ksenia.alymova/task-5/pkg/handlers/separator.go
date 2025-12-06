package handlers

import (
	"context"
	"sync"
)

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return ErrNoChannels
	}

	var wGroup sync.WaitGroup

	wGroup.Add(len(outputs))

	doHandle := func(output chan string) {
		defer wGroup.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case str, ok := <-input:
				if !ok {
					return
				}

				output <- str
			}
		}
	}

	for _, channel := range outputs {
		go doHandle(channel)
	}

	wGroup.Wait()

	return nil
}
