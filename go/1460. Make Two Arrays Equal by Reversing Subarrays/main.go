package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {
	// should be faster than a map as the number of possible values are really small.
	l := [1001]int{}
	for _, elem := range arr {
		l[elem]++
	}
	for _, elem := range target {
		l[elem]--
	}
	for _, elem := range l {
		if elem != 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: target = [1,2,3,4], arr = [2,4,1,3]
	target1 := []int{1, 2, 3, 4}
	arr1 := []int{2, 4, 1, 3}
	fmt.Printf("Test 1: %v\n", canBeEqual(target1, arr1))

	// Test case 2: target = [7], arr = [7]
	target2 := []int{7}
	arr2 := []int{7}
	fmt.Printf("Test 2: %v\n", canBeEqual(target2, arr2))

	// Test case 3: target = [3,7,9], arr = [3,7,11]
	target3 := []int{3, 7, 9}
	arr3 := []int{3, 7, 11}
	fmt.Printf("Test 3: %v\n", canBeEqual(target3, arr3))
}
