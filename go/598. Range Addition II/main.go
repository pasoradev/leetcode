package main

import "math"

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	min_m, min_n := math.MaxInt, math.MaxInt
	for _, op := range ops {
		if op[0] < min_m {
			min_m = op[0]
		}
		if op[1] < min_n {
			min_n = op[1]
		}
	}
	return min_m * min_n
}

func main() {
	result := maxCount(40000, 40000, [][]int{})
	println(result)
}
