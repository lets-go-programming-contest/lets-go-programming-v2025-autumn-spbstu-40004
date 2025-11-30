package handlers

import (
	"context"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	return nil
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	return nil
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	return nil
}
