package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrNoDecorator = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return nil
	}

	var index int

	for {
		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}

			for i := 0; i < len(outputs); i++ {
				select {
				case outputs[index] <- data:
				case <-ctx.Done():
					return ctx.Err()
				default:
				}
				index = (index + 1) % len(outputs)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}

	errCh := make(chan error, len(inputs))
	var wg sync.WaitGroup

	wg.Add(len(inputs))

	for idx := range inputs {
		go func(inputIdx int) {
			defer wg.Done()

			for {
				select {
				case data, ok := <-inputs[inputIdx]:
					if !ok {
						return
					}

					if !strings.Contains(data, "no multiplexer") {
						select {
						case output <- data:
						case <-ctx.Done():
							errCh <- ctx.Err()
							return
						}
					}
				case <-ctx.Done():
					errCh <- ctx.Err()
					return
				}
			}
		}(idx)
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
