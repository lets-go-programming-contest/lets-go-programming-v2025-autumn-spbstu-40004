package handlers

import (
	"context"
)

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return ErrNoChannels
	}

	counterOuput := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case str, ok := <-input:
			if !ok {
				return nil
			}

			outputs[counterOuput%len(outputs)] <- str

			counterOuput++
		}
	}
}
