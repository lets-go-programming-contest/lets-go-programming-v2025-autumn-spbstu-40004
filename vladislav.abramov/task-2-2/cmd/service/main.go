package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var nCount, kCount int

  _, err := fmt.Scan(&nCount)
  if err != nil {
    fmt.Print("Failed to read N\n")
  }

	arr := make([]int, nCount)

	for count := range nCount {
		_, err = fmt.Scan(&arr[count])
    if err != nil {
      fmt.Print("Failed to read data\n")
    }
	}

	_, err = fmt.Scan(&kCount)
  if err != nil {
    fmt.Print("Failed to read K\n")
  }

	heapOfMeals := &IntHeap{}
	heap.Init(heapOfMeals)

	for _, num := range arr {
		heap.Push(heapOfMeals, num)
	}

	var result int

	for range kCount {
		result = heap.Pop(heapOfMeals).(int)
	}

	fmt.Println(result)
}
