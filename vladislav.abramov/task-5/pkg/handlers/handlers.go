package handlers

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	prefix := "decorated: "

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}

			if !strings.HasPrefix(data, prefix) {
				data = prefix + data
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return nil
	}

	var counter int64 = 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			index := atomic.AddInt64(&counter, 1) - 1
			outputIndex := int(index) % len(outputs)

			select {
			case outputs[outputIndex] <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}

	merged := make(chan string, len(inputs)*10)

	for _, input := range inputs {
		go func(in chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}

					select {
					case merged <- data:
					case <-ctx.Done():
						return
					}
				}
			}
		}(input)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-merged:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no multiplexer") {
				continue
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}
