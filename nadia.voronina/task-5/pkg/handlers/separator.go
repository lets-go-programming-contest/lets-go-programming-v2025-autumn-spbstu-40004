package handlers

import (
	"context"
	"fmt"
)

func SeparatorFunc(
	tx context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		return fmt.Errorf("no output channels provided")
	}
	i := 0
	for {
		select {
		case <-tx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}
			out := outputs[i%len(outputs)]
			out <- val
			i++
		}
	}
}
