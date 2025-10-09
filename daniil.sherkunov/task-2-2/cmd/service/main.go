package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initBufSizeBytes = 1 << 16 // 65_536
	maxBufSizeBytes  = 1 << 20 // 1_048_576
)

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {

		return
	}

	*h = append(*h, val)
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]

	return val
}

func (h *MinHeap) Top() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}

	return (*h)[0], true
}

func makeScanner() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, initBufSizeBytes), maxBufSizeBytes)

	return sc
}

func readLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}

	return strings.TrimSpace(scanner.Text()), true
}

func readInt(scanner *bufio.Scanner) (int, bool) {
	text, ok := readLine(scanner)
	if !ok {
		return 0, false
	}

	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, false
	}

	return value, true
}

func readNInts(scanner *bufio.Scanner, expectedCount int) ([]int, bool) {
	collected := make([]int, 0, expectedCount)

	for len(collected) < expectedCount {
		line, ok := readLine(scanner)
		if !ok {
			return nil, false
		}

		for _, tok := range strings.Fields(line) {
			val, err := strconv.Atoi(tok)
			if err != nil {
				return nil, false
			}
			collected = append(collected, val)
			if len(collected) == expectedCount {
				break
			}
		}
	}

	return collected, true
}

func kthLargest(values []int, k int) (int, bool) {
	h := &MinHeap{}
	heap.Init(h)

	for _, v := range values {
		if h.Len() < k {
			heap.Push(h, v)
			continue
		}
		top, _ := h.Top()
		if v > top {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}

	result, ok := h.Top()
	if !ok {
		return 0, false
	}

	return result, true
}

func main() {
	scanner := makeScanner()

	numbersCount, success := readInt(scanner)
	if !success || numbersCount < 1 {
		return
	}

	values, ok := readNInts(scanner, numbersCount)
	if !ok {
		return
	}

	kth, ok2 := readInt(scanner)
	if !ok2 || kth < 1 || kth > numbersCount {
		return
	}

	answer, ok3 := kthLargest(values, kth)
	if !ok3 {
		return
	}

	if _, err := fmt.Println(answer); err != nil {
		_ = err
	}
}
