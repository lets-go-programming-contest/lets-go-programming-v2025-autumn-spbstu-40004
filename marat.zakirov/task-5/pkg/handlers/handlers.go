package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrDecorator = errors.New("this string can't be decorated")

func PrefixDecoratorFunc(cntx context.Context, in chan string, out chan string) error {
	for {
		select {
		case <-cntx.Done():
			return nil
		case wStr, ok := <-in:
			if !ok {
				return nil
			}

			if strings.Contains(wStr, "no multiplexer") {
				return ErrDecorator
			}

			if strings.HasPrefix(wStr, "decorated: ") {
				continue
			}

			select {
			case <-cntx.Done():
				return nil
			case out <- wStr:
			}
		}
	}
}

func SeparatorFunc(cntx context.Context, in chan string, outs []chan string) {
	id := 0

	if len(outs) == 0 {
		return
	}

	for {
		select {
		case <-cntx.Done():
			return
		case wStr, ok := <-in:
			if !ok {
				return
			}
			select {
			case outs[id] <- wStr:
			case <-cntx.Done():
			}
			id = (id + 1) % len(outs)
		}
	}
}

func MultiplexerFunc(cntx context.Context, ins []chan string, out chan string) {
	var wGroup sync.WaitGroup
	multiplex := func(in chan string) {
		defer wGroup.Done()
		for {
			select {
			case <-cntx.Done():
				return
			case wStr, ok := <-in:
				if !ok {
					return
				}

				if strings.Contains(wStr, "no multiplexer") {
					continue
				}

				select {
				case out <- wStr:
				case <-cntx.Done():
					return
				}
			}
		}
	}

	for _, in := range ins {
		wGroup.Add(1)
		go multiplex(in)
	}

	wGroup.Wait()
}
