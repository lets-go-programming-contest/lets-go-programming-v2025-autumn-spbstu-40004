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
			return nil
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
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return nil
	}

	var counter uint64

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			index := atomic.AddUint64(&counter, 1) - 1
			outputIndex := int(index) % len(outputs)

			select {
			case outputs[outputIndex] <- data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}

	type dataWithSource struct {
		data string
	}

	merged := make(chan dataWithSource, len(inputs)*10)

	for i, input := range inputs {
		go func(in chan string, source int) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}
					select {
					case merged <- dataWithSource{data: data}:
					case <-ctx.Done():
						return
					}
				}
			}
		}(input, i)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case item, ok := <-merged:
			if !ok {
				return nil
			}

			if strings.Contains(item.data, "no multiplexer") {
				continue
			}

			select {
			case output <- item.data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}
