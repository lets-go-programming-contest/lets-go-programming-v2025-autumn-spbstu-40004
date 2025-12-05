package handlers

import (
	"context"
)

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			idx := counter % len(outputs)

			select {
			case outputs[idx] <- data:
				counter++
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}
