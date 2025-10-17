package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (maxHeap *MaxHeap) Len() int {
	return len(*maxHeap)
}

func (maxHeap *MaxHeap) Less(indexFirst, indexSecond int) bool {
	return (*maxHeap)[indexFirst] >= (*maxHeap)[indexSecond]
}

func (maxHeap *MaxHeap) Swap(indexFirst, indexSecond int) {
	(*maxHeap)[indexFirst], (*maxHeap)[indexSecond] = (*maxHeap)[indexSecond], (*maxHeap)[indexFirst]
}

func (maxHeap *MaxHeap) Push(value any) {
	val, ok := value.(int)
	if !ok {
		panic("Value must be int in MaxHeap")
	}

	*maxHeap = append(*maxHeap, val)
}

func (maxHeap *MaxHeap) Pop() any {
	oldLen := len(*maxHeap)
	if oldLen == 0 {
		return nil
	}

	oldHeap := *maxHeap
	poppedValue := oldHeap[oldLen - 1]
	*maxHeap = oldHeap[0 : oldLen - 1]

	return poppedValue
}

func main() {
	var (
		amount     uint
		priority   int
		preference uint
	)

	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Error while reading amount of meals")

		return
	}

	preferences := &maxheap.MaxHeap{}
	heap.Init(preferences)

	for range amount {
		_, fmt.Scan(&priority)
		if err != nil {
			fmt.Println("Error while reading meal priority")

			return
		}

		heap.Push(preferences, priority)
	}

	_, err = fmt.Scan(&preference)
	if err != nil {
		fmt.Println("Error while reading meal preference")

		return
	}

	for range preference - 1 {
		heap.Pop(preferences)
	}

	result, ok := heap.Pop(preferences).(int)
	if ok {
		fmt.Println(result)
	}
}
