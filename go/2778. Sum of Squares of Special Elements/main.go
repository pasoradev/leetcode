package main

import "fmt"

func sumOfSquares(nums []int) int {
	l := len(nums)
	total := 0
	for i := 1; i <= l; i++ {
		if l%i == 0 {
			total += nums[i-1] * nums[i-1]
		}
	}
	return total
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 3, 4}
	result1 := sumOfSquares(nums1)
	fmt.Printf("Example 1: %d (expected 21)\n", result1)

	// Example 2
	nums2 := []int{2, 7, 1, 19, 18, 3}
	result2 := sumOfSquares(nums2)
	fmt.Printf("Example 2: %d (expected 63)\n", result2)
}
