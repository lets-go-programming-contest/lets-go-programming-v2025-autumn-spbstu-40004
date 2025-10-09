package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any          { old := *h; v := old[len(old)-1]; *h = old[:len(old)-1]; return v }
func (h MinHeap) Top() (int, bool) {
	if len(h) == 0 {
		return 0, false
	}
	return h[0], true
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 0, 1<<16), 1<<20)

	readLine := func() (string, bool) {
		if !in.Scan() {
			return "", false
		}
		return strings.TrimSpace(in.Text()), true
	}
	readInt := func() (int, bool) {
		s, ok := readLine()
		if !ok {
			return 0, false
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0, false
		}
		return v, true
	}

	N, ok := readInt()
	if !ok || N < 1 {
		return
	}

	line, ok := readLine()
	if !ok {
		return
	}
	tokens := strings.Fields(line)
	if len(tokens) != N {
		for len(tokens) < N && in.Scan() {
			tokens = append(tokens, strings.Fields(in.Text())...)
		}
		if len(tokens) != N {
			return
		}
	}
	values := make([]int, 0, N)
	for _, t := range tokens {
		v, err := strconv.Atoi(t)
		if err != nil {
			return
		}
		values = append(values, v)
	}

	k, ok := readInt()
	if !ok || k < 1 || k > N {
		return
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, v := range values {
		if h.Len() < k {
			heap.Push(h, v)
		} else if v > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}
	ans, _ := h.Top()
	fmt.Println(ans)
}
