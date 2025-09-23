package main

import (
	"fmt"
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	min := math.MaxInt
	res := []string{}

	for i, s := range list1 {
		m[s] = i
	}

	for i, s := range list2 {
		_, exists := m[s]
		if exists {
			m[s] += i
			if m[s] < min {
				min = m[s]
				res = []string{s}
			} else if min == m[s] {
				res = append(res, s)
			}
		}
	}
	return res
}

func main() {
	result := findRestaurant(
		[]string{"happy", "sad", "good"},
		[]string{"sad", "happy", "good"},
	)
	fmt.Println(result)
}
