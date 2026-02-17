package main

import (
	"fmt"
	"slices"
)

func canMakeEqual(nums []int, k int) bool {
	numsCopy := slices.Clone(nums)
	kCopy := k
	satisfies := true
	// try to make everything negative
	for i := 0; i < len(numsCopy)-1; i++ {
		num := numsCopy[i]
		if kCopy < 0 {
			satisfies = false
			break
		}

		if num == -1 {
			kCopy--
			numsCopy[i+1] *= -1
		}
	}

	if satisfies && kCopy >= 0 && numsCopy[len(numsCopy)-1] == 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		num := nums[i]
		if k < 0 {
			return false
		}

		if num == 1 {
			k--
			nums[i+1] *= -1
		}
	}
	return k >= 0 && nums[len(nums)-1] == -1
}

func main() {
	nums1 := []int{1, -1, 1, -1, 1}
	k1 := 3
	result1 := canMakeEqual(nums1, k1)
	fmt.Printf("Example 1: %v (expected true)\n", result1)

	nums2 := []int{-1, -1, -1, 1, 1, 1}
	k2 := 5
	result2 := canMakeEqual(nums2, k2)
	fmt.Printf("Example 2: %v (expected false)\n", result2)

	nums3 := []int{1, -1, 1}
	k3 := 2
	result3 := canMakeEqual(nums3, k3)
	fmt.Printf("Example 3: %v (expected true)\n", result3)

}
