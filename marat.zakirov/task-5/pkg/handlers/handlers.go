package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrDecorator = errors.New("this string can't be decorated")

func PrefixDecoratorFunc(cntx context.Context, inChannelP chan string, outChannelP chan string) error {
	for {
		select {
		case <-cntx.Done():
			return nil
		case wStr, ok := <-inChannelP:
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
			case outChannelP <- wStr:
			}
		}
	}
}

func SeparatorFunc(cntx context.Context, inChannelP chan string, outChannelsP []chan string) {
	ChannelID := 0

	if len(outChannelsP) == 0 {
		return
	}

	for {
		select {
		case <-cntx.Done():
			return
		case wStr, ok := <-inChannelP:
			if !ok {
				return
			}
			select {
			case outChannelsP[ChannelID] <- wStr:
			case <-cntx.Done():
			}

			ChannelID = (ChannelID + 1) % len(outChannelsP)
		}
	}
}

func MultiplexerFunc(cntx context.Context, inChannelsP []chan string, outChannelP chan string) {
	var wGroup sync.WaitGroup

	multiplex := func(inputChan chan string) {
		defer wGroup.Done()

		for {
			select {
			case <-cntx.Done():
				return
			case wStr, ok := <-inputChan:
				if !ok {
					return
				}

				if strings.Contains(wStr, "no multiplexer") {
					continue
				}

				select {
				case outChannelP <- wStr:
				case <-cntx.Done():
					return
				}
			}
		}
	}

	for _, in := range inChannelsP {
		wGroup.Add(1)

		go multiplex(in)
	}

	wGroup.Wait()
}
