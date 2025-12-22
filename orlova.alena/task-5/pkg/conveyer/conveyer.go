package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrChan = errors.New("chan not found")

type ConveyerImpl struct {
	mu           sync.RWMutex
	size         int
	channels     map[string]chan string
	decorators   []decoratorSpec
	multiplexers []multiplexerSpec
	separators   []separatorSpec
}

type decoratorSpec struct {
	fn     func(ctx context.Context, input chan string, output chan string) error
	input  string
	output string
}

type separatorSpec struct {
	fn      func(ctx context.Context, input chan string, outputs []chan string) error
	input   string
	outputs []string
}

type multiplexerSpec struct {
	fn     func(ctx context.Context, inputs []chan string, output chan string) error
	inputs []string
	output string
}

func New(size int) *ConveyerImpl {
	return &ConveyerImpl{
		mu:           sync.RWMutex{},
		size:         size,
		channels:     make(map[string]chan string),
		decorators:   make([]decoratorSpec, 0),
		multiplexers: make([]multiplexerSpec, 0),
		separators:   make([]separatorSpec, 0),
	}
}

func (conv *ConveyerImpl) createChannel(id string) {
	conv.mu.Lock()
	defer conv.mu.Unlock()

	if _, exists := conv.channels[id]; !exists {
		conv.channels[id] = make(chan string, conv.size)
	}
}

func (conv *ConveyerImpl) getChannel(id string) (chan string, error) {
	conv.mu.RLock()
	defer conv.mu.RUnlock()

	ch, exists := conv.channels[id]
	if !exists {
		return nil, ErrChan
	}

	return ch, nil
}

func (conv *ConveyerImpl) closeChannels() {
	conv.mu.Lock()
	defer conv.mu.Unlock()

	for _, channel := range conv.channels {
		close(channel)
	}
}

func (conv *ConveyerImpl) RegisterDecorator(
	decFn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	conv.createChannel(input)
	conv.createChannel(output)

	conv.mu.Lock()
	conv.decorators = append(conv.decorators, decoratorSpec{
		fn:     decFn,
		input:  input,
		output: output,
	})
	conv.mu.Unlock()
}

func (conv *ConveyerImpl) RegisterMultiplexer(
	mulFn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inp := range inputs {
		conv.createChannel(inp)
	}

	conv.createChannel(output)

	conv.mu.Lock()
	conv.multiplexers = append(conv.multiplexers, multiplexerSpec{
		fn:     mulFn,
		inputs: inputs,
		output: output,
	})
	conv.mu.Unlock()
}

func (conv *ConveyerImpl) RegisterSeparator(
	regFn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	conv.createChannel(input)

	for _, out := range outputs {
		conv.createChannel(out)
	}

	conv.mu.Lock()
	conv.separators = append(conv.separators, separatorSpec{
		fn:      regFn,
		input:   input,
		outputs: outputs,
	})
	conv.mu.Unlock()
}

func (conv *ConveyerImpl) Run(ctx context.Context) error {
	group, groupCtx := errgroup.WithContext(ctx)

	for _, decorator := range conv.decorators {
		dec := decorator

		group.Go(func() error {
			input, _ := conv.getChannel(dec.input)
			output, _ := conv.getChannel(dec.output)

			return dec.fn(groupCtx, input, output)
		})
	}

	for _, multiplexer := range conv.multiplexers {
		mul := multiplexer

		group.Go(func() error {
			inputs := make([]chan string, len(mul.inputs))

			for i, name := range mul.inputs {
				inputs[i], _ = conv.getChannel(name)
			}

			output, _ := conv.getChannel(mul.output)

			return mul.fn(groupCtx, inputs, output)
		})
	}

	for _, separator := range conv.separators {
		sep := separator

		group.Go(func() error {
			input, _ := conv.getChannel(sep.input)
			outputs := make([]chan string, len(sep.outputs))

			for i, name := range sep.outputs {
				outputs[i], _ = conv.getChannel(name)
			}

			return sep.fn(groupCtx, input, outputs)
		})
	}

	err := group.Wait()

	conv.closeChannels()

	if err != nil {
		return fmt.Errorf("conveyer error: %w", err)
	}

	return nil
}

func (conv *ConveyerImpl) Send(input string, data string) error {
	ch, err := conv.getChannel(input)
	if err != nil {
		return err
	}

	ch <- data

	return nil
}

func (conv *ConveyerImpl) Recv(output string) (string, error) {
	ch, err := conv.getChannel(output)
	if err != nil {
		return "", err
	}

	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return data, nil
}
