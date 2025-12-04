package handlers

import (
	"context"
	"strings"
)

func MultiplexerFunc(
	tx context.Context,
	inputs []chan string,
	output chan string,
) error {
	closedCount := 0
	totalInputs := len(inputs)
	closed := make([]bool, totalInputs)

	for {
		select {
		case <-tx.Done():
			return tx.Err()
		default:
			for i, in := range inputs {
				if closed[i] {
					continue
				}
				select {
				case msg, ok := <-in:
					if !ok {
						closed[i] = true
						closedCount++
						continue
					}
					if strings.Contains(msg, "no multiplexer") {
						continue
					}
					select {
					case output <- msg:
					case <-tx.Done():
						return tx.Err()
					}
				default:
				}
			}
			if closedCount == totalInputs {
				return nil
			}
		}
	}
}
