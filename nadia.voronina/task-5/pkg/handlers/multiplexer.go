package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(
	cntxt context.Context,
	inputs []chan string,
	output chan string,
) error {
	var waitGroup sync.WaitGroup

	totalInputs := len(inputs)

	waitGroup.Add(totalInputs)

	for _, input := range inputs {
		go func(channel chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-cntxt.Done():
					return
				case msg, ok := <-channel:
					if !ok {
						return
					}

					if strings.Contains(msg, "no multiplexer") {
						continue
					}
					select {
					case output <- msg:
					case <-cntxt.Done():
						return
					}
				}
			}
		}(input)
	}

	waitGroup.Wait()

	return nil
}
