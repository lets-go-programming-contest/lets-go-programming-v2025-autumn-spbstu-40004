package handlers

import (
	"context"
	"strings"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	const skipPhrase = "no multiplexer"

	done := make(chan struct{})
	defer close(done)

	for _, input := range inputs {
		go func(in chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}
					if strings.Contains(data, skipPhrase) {
						continue
					}
					select {
					case output <- data:
					case <-ctx.Done():
						return
					case <-done:
						return
					}
				}
			}
		}(input)
	}

	<-ctx.Done()
	return ctx.Err()
}
