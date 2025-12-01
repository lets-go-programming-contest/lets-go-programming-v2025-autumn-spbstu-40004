package handlers_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/15446-rus75/task-5/pkg/handlers"
)

func TestPrefixDecoratorFunc_Basic(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan string, 1)
	output := make(chan string, 1)

	input <- "test"
	close(input)

	go func() {
		err := handlers.PrefixDecoratorFunc(ctx, input, output)
		assert.NoError(t, err)
	}()

	result := <-output
	assert.Equal(t, "decorated: test", result)
}

func TestPrefixDecoratorFunc_AlreadyDecorated(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan string, 1)
	output := make(chan string, 1)

	input <- "decorated: already"
	close(input)

	go func() {
		err := handlers.PrefixDecoratorFunc(ctx, input, output)
		require.NoError(t, err)
	}()

	result := <-output
	assert.Equal(t, "decorated: already", result)
}

func TestPrefixDecoratorFunc_Error(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan string, 1)
	output := make(chan string, 1)

	input <- "text with no decorator"
	close(input)

	err := handlers.PrefixDecoratorFunc(ctx, input, output)
	assert.ErrorIs(t, err, handlers.ErrNoDecorator)
}

func TestPrefixDecoratorFunc_ContextCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	input := make(chan string)
	output := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	err := handlers.PrefixDecoratorFunc(ctx, input, output)
	assert.NoError(t, err)
}

func TestSeparatorFunc_RoundRobin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan string, 3)
	outputs := []chan string{
		make(chan string, 2),
		make(chan string, 2),
	}

	input <- "a"
	input <- "b"
	input <- "c"
	close(input)

	go func() {
		err := handlers.SeparatorFunc(ctx, input, outputs)
		require.NoError(t, err)
	}()

	assert.Equal(t, "a", <-outputs[0])
	assert.Equal(t, "b", <-outputs[1])
	assert.Equal(t, "c", <-outputs[0])
}

func TestSeparatorFunc_NoOutputs(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan string, 1)
	input <- "test"
	close(input)

	err := handlers.SeparatorFunc(ctx, input, []chan string{})
	assert.NoError(t, err)
}

func TestSeparatorFunc_SingleOutput(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan string, 2)
	output := make(chan string, 2)

	input <- "first"
	input <- "second"
	close(input)

	go func() {
		err := handlers.SeparatorFunc(ctx, input, []chan string{output})
		require.NoError(t, err)
	}()

	assert.Equal(t, "first", <-output)
	assert.Equal(t, "second", <-output)
}

func TestMultiplexerFunc_Basic(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputs := []chan string{
		make(chan string, 1),
		make(chan string, 1),
	}
	output := make(chan string, 2)

	inputs[0] <- "from1"
	inputs[1] <- "from2"
	close(inputs[0])
	close(inputs[1])

	go func() {
		err := handlers.MultiplexerFunc(ctx, inputs, output)
		assert.NoError(t, err)
	}()

	results := []string{<-output, <-output}
	assert.ElementsMatch(t, []string{"from1", "from2"}, results)
}

func TestMultiplexerFunc_Filter(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputs := []chan string{make(chan string, 2)}
	output := make(chan string, 2)

	inputs[0] <- "normal"
	inputs[0] <- "no multiplexer here"
	close(inputs[0])

	go func() {
		err := handlers.MultiplexerFunc(ctx, inputs, output)
		require.NoError(t, err)
	}()

	result := <-output
	assert.Equal(t, "normal", result)

	select {
	case <-output:
		t.Fatal("should not receive filtered message")
	case <-time.After(10 * time.Millisecond):
	}
}

func TestMultiplexerFunc_NoInputs(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	output := make(chan string)

	err := handlers.MultiplexerFunc(ctx, []chan string{}, output)
	assert.NoError(t, err)
}
