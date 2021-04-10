package main

import (
	"fmt"
)

var blocks [26]int

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinimumBlocks(m, n int) int {
	result := make([]int, m+1)
	result[0] = 0
	for i := 1; i <= m; i++ {
		result[i] = 1000001
	}

	for j := 0; j < n; j++ {
		for i := blocks[j]; i <= m; i++ {
			result[i] = Min(result[i], result[i-blocks[j]]+1)
		}
	}
	return result[m]
}

func InitBlocks() {
	for i := 0; i < 25; i++ {
		blocks[i] = 0
	}
}

func main() {
	var t, n, m int
	fmt.Scanf("%d", &t)

	for j := 0; j < t; j++ {
		fmt.Scanf("%d %d", &n, &m)

		InitBlocks()

		for i := 0; i < n; i++ {
			fmt.Scanln(&blocks[i])
		}

		fmt.Printf("%d\n", MinimumBlocks(m, n))
	}
}
