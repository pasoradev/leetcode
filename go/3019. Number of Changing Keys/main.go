package main

import "fmt"

func countKeyChanges(s string) int {
	count := 0
	last := rune(s[0])
	for _, c := range s {
		if !(c == last || c+32 == last || c-32 == last) {
			last = c
			count++
		}
	}
	return count
}

func main() {
	// Example 1
	s1 := "aAbBcC"
	result1 := countKeyChanges(s1)
	fmt.Printf("Example 1: %d (expected 2)\n", result1)

	// Example 2
	s2 := "AaAaAaaA"
	result2 := countKeyChanges(s2)
	fmt.Printf("Example 2: %d (expected 0)\n", result2)
}
