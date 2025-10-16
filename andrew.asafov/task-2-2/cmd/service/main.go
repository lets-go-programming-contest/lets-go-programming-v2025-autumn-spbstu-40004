package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&k)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	result := arr[k-1]
	fmt.Println(result)
}
