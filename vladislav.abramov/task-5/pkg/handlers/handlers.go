package handlers

import (
	"context"
	"errors"
	"fmt"
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
				return fmt.Errorf("%w", ErrNoDecorator)
			}

			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return fmt.Errorf("%w", ctx.Err())
			}
		case <-ctx.Done():
			return fmt.Errorf("%w", ctx.Err())
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

			for range outputs {
				select {
				case outputs[index] <- data:
				case <-ctx.Done():
					return fmt.Errorf("%w", ctx.Err())
				default:
				}
				index = (index + 1) % len(outputs)
			}
		case <-ctx.Done():
			return fmt.Errorf("%w", ctx.Err())
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return nil
	}

	errCh := make(chan error, len(inputs))
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(inputs))

	for idx := range inputs {
		go func(inputIdx int) {
			defer waitGroup.Done()

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
							errCh <- fmt.Errorf("%w", ctx.Err())
							return
						}
					}
				case <-ctx.Done():
					errCh <- fmt.Errorf("%w", ctx.Err())
					return
				}
			}
		}(idx)
	}

	waitGroup.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
