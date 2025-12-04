package handlers

import (
	"context"
)

func SeparatorFunc(
	tx context.Context,
	input chan string,
	outputs []chan string,
) error {
	i := 0
	for {
		select {
		case <-tx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}
			if len(outputs) == 0 {
				continue
			}
			out := outputs[i%len(outputs)]
			out <- val
			i++
		}
	}
}
