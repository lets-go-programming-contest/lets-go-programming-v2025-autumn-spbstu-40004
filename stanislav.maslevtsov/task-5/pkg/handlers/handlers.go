package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrNoDecorator = errors.New("can't be decorated")
	ErrNoChans     = errors.New("no channels received")
)

func PrefixDecoratorFunc(
	ctx context.Context,
	input, output chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case str, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(str, "no decorator") {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(str, "decorated: ") {
				str = "decorated: " + str
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- str:
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	if len(inputs) == 0 {
		return ErrNoChans
	}

	var waitGr sync.WaitGroup

	waitGr.Add(len(inputs))

	for ichanIdx := range inputs {
		go func() {
			defer waitGr.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case str, ok := <-inputs[ichanIdx]:
					if !ok {
						return
					}

					if !strings.Contains(str, "no multiplexer") {
						select {
						case <-ctx.Done():
							return
						case output <- str:
						}
					}
				}
			}
		}()
	}

	waitGr.Wait()

	return nil
}

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		return ErrNoChans
	}

	var ochanIdx int

	for {
		select {
		case <-ctx.Done():
			return nil
		case str, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return nil
			case outputs[ochanIdx] <- str:
			}

			ochanIdx = (ochanIdx + 1) % len(outputs)
		}
	}
}
