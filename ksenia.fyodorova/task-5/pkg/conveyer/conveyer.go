package conveyer

import (
	"context"
	"sync"
)

type DecoratorFunc func(ctx context.Context, input chan string, output chan string) error
type MultiplexerFunc func(ctx context.Context, inputs []chan string, output chan string) error
type SeparatorFunc func(ctx context.Context, input chan string, outputs []chan string) error

type Conveyer interface {
	RegisterDecorator(fn DecoratorFunc, input string, output string)
	RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string)
	RegisterSeparator(fn SeparatorFunc, input string, outputs []string)
	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

type ConveyerImpl struct {
	chanManager *ChanManager
	size        int
	handlers    []func(context.Context) error
	mu          sync.Mutex
}

func New(size int) *ConveyerImpl {
	return &ConveyerImpl{
		chanManager: NewChanManager(),
		size:        size,
	}
}

func (c *ConveyerImpl) RegisterDecorator(fn DecoratorFunc, input string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		inChan := c.chanManager.GetOrCreate(input, c.size)
		outChan := c.chanManager.GetOrCreate(output, c.size)
		return fn(ctx, inChan, outChan)
	})
}

func (c *ConveyerImpl) RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		inChans := make([]chan string, len(inputs))
		for i, input := range inputs {
			inChans[i] = c.chanManager.GetOrCreate(input, c.size)
		}
		outChan := c.chanManager.GetOrCreate(output, c.size)
		return fn(ctx, inChans, outChan)
	})
}

func (c *ConveyerImpl) RegisterSeparator(fn SeparatorFunc, input string, outputs []string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers = append(c.handlers, func(ctx context.Context) error {
		inChan := c.chanManager.GetOrCreate(input, c.size)
		outChans := make([]chan string, len(outputs))
		for i, output := range outputs {
			outChans[i] = c.chanManager.GetOrCreate(output, c.size)
		}
		return fn(ctx, inChan, outChans)
	})
}

func (c *ConveyerImpl) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	errCh := make(chan error, len(c.handlers))

	c.mu.Lock()
	handlers := make([]func(context.Context) error, len(c.handlers))
	copy(handlers, c.handlers)
	c.mu.Unlock()

	for _, handler := range handlers {
		wg.Add(1)
		go func(h func(context.Context) error) {
			defer wg.Done()
			if err := h(ctx); err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(handler)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err, ok := <-errCh:
		if ok {
			cancel()
			wg.Wait()
			c.chanManager.CloseAll()
			return err
		}
	case <-ctx.Done():
		wg.Wait()
		c.chanManager.CloseAll()
		return ctx.Err()
	}

	wg.Wait()
	c.chanManager.CloseAll()
	return nil
}

func (c *ConveyerImpl) Send(input string, data string) error {
	return c.chanManager.Send(input, data)
}

func (c *ConveyerImpl) Recv(output string) (string, error) {
	return c.chanManager.Recv(output)
}
