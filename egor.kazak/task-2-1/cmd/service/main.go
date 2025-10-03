package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	for i := 0; i < N; i++ {
		const min_T = 15
		min_t := 15
		const max_T = 30
		max_t := 30

		for j := 0; j < K; j++ {
			var op string
			var t int
			fmt.Scan(&op, &t)

			switch op {
			case ">=":
				if t > min_t {
					min_t = t
				}
			case "<=":
				if t < max_t {
					max_t = t
				}
			}

			if min_t <= max_t {
				fmt.Println(min_t)
			} else if min_t >= max_t || t <= min_T || t >= max_T {
				fmt.Println(-1)
			}
		}
	}
}
