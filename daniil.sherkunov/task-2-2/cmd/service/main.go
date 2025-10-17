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
	initBufSizeBytes = 1 << 16
	maxBufSizeBytes  = 1 << 20
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

	buf := append(*h, val)
	*h = buf
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
	text, textOK := readLine(scanner)
	if !textOK {
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
		line, lineOK := readLine(scanner)
		if !lineOK {
			return nil, false
		}

		tokens := strings.Fields(line)

		for _, token := range tokens {
			val, err := strconv.Atoi(token)
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

func kthLargest(values []int, kth int) (int, bool) {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, value := range values {
		if minHeap.Len() < kth {
			heap.Push(minHeap, value)

			continue
		}

		top, _ := minHeap.Top()

		if value > top {
			heap.Pop(minHeap)
			heap.Push(minHeap, value)
		}
	}

	result, ok := minHeap.Top()
	if !ok {
		return 0, false
	}

	return result, true
}

func main() {
	scanner := makeScanner()

	numbersCount, numbersCountOK := readInt(scanner)
	if !numbersCountOK || numbersCount < 1 {
		return
	}

	values, valuesOK := readNInts(scanner, numbersCount)
	if !valuesOK {
		return
	}

	kth, kthOK := readInt(scanner)
	if !kthOK || kth < 1 || kth > numbersCount {
		return
	}

	answer, answerOK := kthLargest(values, kth)
	if !answerOK {
		return
	}

	if _, err := fmt.Println(answer); err != nil {
		_ = err
	}
}
