package intmaxheap

type IntMaxHeap []int

func (h *IntMaxHeap) Len() int {
	return len(*h)
}

func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(val any) {
	intVal, isCorrect := val.(int)
	if isCorrect {
		*h = append(*h, intVal)
	} else {
		panic("Error: incorrect input")
	}
}

func (h *IntMaxHeap) Pop() any {
	prev := *h
	prevLen := len(prev)

	if prevLen == 0 {
		return nil
	}

	el := prev[prevLen-1]
	*h = prev[:prevLen-1]

	return el
}
