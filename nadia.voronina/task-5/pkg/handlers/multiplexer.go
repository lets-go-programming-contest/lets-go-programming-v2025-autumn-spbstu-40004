package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(
	tx context.Context,
	inputs []chan string,
	output chan string,
) error {
	var wg sync.WaitGroup
	totalInputs := len(inputs)

	wg.Add(totalInputs)
	for _, in := range inputs {
		go func(ch chan string) {
			defer wg.Done()
			for {
				select {
				case <-tx.Done():
					return
				case msg, ok := <-ch:
					if !ok {
						return
					}
					if strings.Contains(msg, "no multiplexer") {
						continue
					}
					select {
					case output <- msg:
					case <-tx.Done():
						return
					}
				}
			}
		}(in)
	}

	wg.Wait()
	return nil
}
