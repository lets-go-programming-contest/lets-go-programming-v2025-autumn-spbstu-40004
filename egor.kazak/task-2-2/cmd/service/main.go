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
		log.Fatal("invalid preference order index")
	}

	h := &heaputil.IntHeap{}
	heaputil.Init(h)

	for _, rating := range ratings {
		heaputil.Push(h, rating)
	}

	var result int
	for i := 0; i < preferenceOrder; i++ {
		result = heaputil.Pop(h)
	}

	fmt.Println(result)
}