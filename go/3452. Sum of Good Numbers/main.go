package main

import "fmt"

func isGood(nums []int, i, num, k int) bool {
	if i-k >= 0 && num <= nums[i-k] {
		return false
	}
	if i+k < len(nums) && num <= nums[i+k] {
		return false
	}
	return true
}

func sumOfGoodNumbers(nums []int, k int) int {
	s := 0
	for i, num := range nums {
		if isGood(nums, i, num, k) {
			s += num
		}
	}
	return s
}

func main() {
	// Example 1
	nums1 := []int{1, 3, 2, 1, 5, 4}
	k1 := 2
	result1 := sumOfGoodNumbers(nums1, k1)
	fmt.Printf("Example 1: %d (expected 12)\n", result1)

	// Example 2
	nums2 := []int{2, 1}
	k2 := 1
	result2 := sumOfGoodNumbers(nums2, k2)
	fmt.Printf("Example 2: %d (expected 2)\n", result2)
}
