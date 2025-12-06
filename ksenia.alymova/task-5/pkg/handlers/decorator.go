package handlers

import (
	"context"
	"errors"
	"strings"
)

const (
	undecorated   = "no decorator"
	decoratedPref = "decorated: "
)

var ErrUndecorated = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case str, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(str, undecorated) {
				return ErrUndecorated
			}

			newStr := str
			if !strings.HasPrefix(str, decoratedPref) {
				newStr = decoratedPref + str
			}

			output <- newStr
		}
	}
}
