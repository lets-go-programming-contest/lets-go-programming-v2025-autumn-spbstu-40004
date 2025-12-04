package main

import (
	"context"
	"time"

	"spbstu.ru/nadia.voronina/task-5/pkg/conveyer"
	"spbstu.ru/nadia.voronina/task-5/pkg/handlers"
)

func main() {
	cnv := conveyer.New(10)
	cnv.RegisterDecorator(handlers.PrefixDecoratorFunc, "input1", "decOutput1")
	cnv.RegisterDecorator(handlers.PrefixDecoratorFunc, "input2", "decOutput2")
	cnv.RegisterMultiplexer(handlers.MultiplexerFunc, []string{"decOutput1", "decOutput2"}, "muxOutput")
	cnv.RegisterSeparator(handlers.SeparatorFunc, "muxOutput", []string{"output1", "output2"})
	ctx := context.Background()
	go cnv.Run(ctx)

	cnv.Send("input1", "data1")
	cnv.Send("input2", "data2")
	cnv.Send("input2", "no decorator")

	time.Sleep(1 * time.Second)
	out1, _ := cnv.Recv("output1")
	out2, _ := cnv.Recv("output2")
	println("Output1:", out1)
	println("Output2:", out2)
}
