package intheap

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	value, err := x.(int)
	if !err {
		panic("Invalid type")
	} else {
		*h = append(*h, value)
	}
}

func (h *IntHeap) Pop() any {
	old := *h
	count := len(old)

	if count == 0 {
		panic("Pop from empty heap")
	}

	x := old[count-1]
	*h = old[0 : count-1]

	return x
}
