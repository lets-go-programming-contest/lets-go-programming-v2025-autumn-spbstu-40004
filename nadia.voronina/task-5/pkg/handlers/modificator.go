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
			if s == "" {
				continue
			}
			if strings.Contains(s, "no decorator") {
				return fmt.Errorf("can't be decorated: %s", s)
			}
			prefix := "decorated: "
			if strings.HasPrefix(s, prefix) {
				output <- s
			} else {
				output <- prefix + s
			}
		}
	}
}
