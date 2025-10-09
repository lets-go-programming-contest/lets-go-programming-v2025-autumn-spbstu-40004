package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 0, 1<<16), 1<<20)

	readInt := func() (int, bool) {
		if !in.Scan() {
			return 0, false
		}
		s := strings.TrimSpace(in.Text())
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

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for dept := 0; dept < N; dept++ {
		K, ok := readInt()
		if !ok || K < 1 {
			return
		}
		lo, hi := minTemp, maxTemp

		for i := 0; i < K; i++ {
			if !in.Scan() {
				return
			}
			line := strings.TrimSpace(in.Text())
			fields := strings.Fields(line)
			if len(fields) != 2 {
				return
			}
			op := fields[0]
			val, err := strconv.Atoi(fields[1])
			if err != nil {
				return
			}

			switch op {
			case ">=", "≥":
				if val > lo {
					lo = val
				}
			case "<=", "≤":
				if val < hi {
					hi = val
				}
			default:
				return
			}

			if lo > hi {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, lo)
			}
		}
	}
}
