package handlers

import (
	"context"
	"fmt"
	"strings"
)

type DecoratorError struct {
	Msg string
}

func (e *DecoratorError) Error() string {
	return fmt.Sprintf("can't be decorated: %s", e.Msg)
}

func PrefixDecoratorFunc(
	tx context.Context,
	input chan string,
	output chan string,
) error {
	for {
		select {
		case <-tx.Done():
			return nil
		case s, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(s, "no decorator") {
				return &DecoratorError{Msg: s}
			}

			prefix := "decorated: "
			select {
			case <-tx.Done():
				return nil
			case output <- func() string {
				if strings.HasPrefix(s, prefix) {
					return s
				}
				return prefix + s
			}():
			}
		}
	}
}
