package handlers

import (
	"context"
	"strings"
)

type DecoratorError struct {
	Msg string
}

func (e *DecoratorError) Error() string {
	return "can't be decorated " + e.Msg
}

func PrefixDecoratorFunc(
	cntxt context.Context,
	input chan string,
	output chan string,
) error {
	for {
		select {
		case <-cntxt.Done():
			return nil
		case incomingStr, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(incomingStr, "no decorator") {
				return &DecoratorError{Msg: incomingStr}
			}

			prefix := "decorated: "
			select {
			case <-cntxt.Done():
				return nil
			case output <- func() string {
				if strings.HasPrefix(incomingStr, prefix) {
					return incomingStr
				}

				return prefix + incomingStr
			}():
			}
		}
	}
}
