package main

import "fmt"

func reverse(r []rune, l int) {
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func finalString(s string) string {
	r := []rune(s)
	iCount := 0
	for i := range r {
		if r[i] == 'i' {
			iCount++
			reverse(r, i)
		}
	}

	res := make([]rune, len(s)-iCount)
	j := 0
	for i := range r {
		if r[i] != 'i' {
			res[j] = r[i]
			j++
		}
	}
	return string(res)
}

func main() {
	// Example 1
	s1 := "string"
	result1 := finalString(s1)
	fmt.Printf("Example 1: %q (expected \"rtsng\")\n", result1)

	// Example 2
	s2 := "poiinter"
	result2 := finalString(s2)
	fmt.Printf("Example 2: %q (expected \"ponter\")\n", result2)
}
