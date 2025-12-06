package handlers

import (
	"context"
	"strings"
	"sync"
)

const unmultiplexered = "no multiplexer"

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var wGroup sync.WaitGroup

	wGroup.Add(len(inputs))

	doHandle := func(input chan string) {
		defer wGroup.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case str, ok := <-input:
				if !ok {
					return
				}

				if !strings.Contains(str, unmultiplexered) {
					output <- str
				}
			}
		}
	}

	for _, channel := range inputs {
		go doHandle(channel)
	}

	wGroup.Wait()

	return nil
}
