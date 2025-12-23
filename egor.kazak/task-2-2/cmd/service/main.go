package main

import (
	"fmt"
	"log"

	"github.com/CuatHimBong/task-2-2/internal/heaputil"
)

func main() {
	var dishCount, preferenceOrder int

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		log.Fatal(err)
	}

	ratings := make([]int, dishCount)
	for i := range ratings {
		_, err := fmt.Scan(&ratings[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		log.Fatal(err)
	}

	if preferenceOrder <= 0 || preferenceOrder > dishCount {
		log.Fatal("preferenceOrder must be in range [1, dishCount]")
	}

	heap := &heaputil.IntHeap{}
	heaputil.Init(heap)

	for _, rating := range ratings {
		heaputil.Push(heap, rating)
	}

	var result int
	for range preferenceOrder {
		result = heaputil.Pop(heap)
	}

	fmt.Println(result)
}
