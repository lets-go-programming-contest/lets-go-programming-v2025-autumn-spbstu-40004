package handlers

import (
	"context"
	"fmt"
	"strings"
)

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
				return fmt.Errorf("can't be decorated: %s", s)
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
