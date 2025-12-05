package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(inputs))

	for index := range inputs {
		go func(idx int) {
			defer waitGroup.Done()

			for {
				select {
				case data, ok := <-inputs[idx]:
					if !ok {
						return
					}

					if !strings.Contains(data, "no multiplexer") {
						select {
						case output <- data:
						case <-ctx.Done():
							return
						}
					}
				case <-ctx.Done():
					return
				}
			}
		}(index)
	}

	waitGroup.Wait()

	return nil
}
