package main

import (
	"fmt"
	"slices"
)

type ApplesPack struct {
	amount int
	expiry int
}

func eat(l []ApplesPack, currentDay int, eatenApplesAmount *int) {
	length := len(l)
	if length != 0 && l[length-1].amount > 0 && l[length-1].expiry > currentDay {
		(*eatenApplesAmount)++
		l[length-1].amount--
	}
}

func cleanup(l []ApplesPack, currentDay int) []ApplesPack {
	for i := len(l) - 1; i >= 0; i-- {
		pack := l[i]
		if currentDay >= pack.expiry || pack.amount <= 0 {
			l = slices.Delete(l, i, i+1)
		}
	}
	return l
}

func eatenApples(apples []int, days []int) int {
	eatenApplesAmount := 0
	l := make([]ApplesPack, 0)

	currentDay := 0

	for _, daysGoodFor := range days {
		inserted := false
		expiry := currentDay + daysGoodFor
		// try to insert into before the next longer expiry apple batch
		for i := len(l) - 1; i >= 0; i-- {
			pack := l[i]
			if pack.expiry > expiry {
				inserted = true
				l = slices.Insert(l, i+1, ApplesPack{amount: apples[currentDay], expiry: expiry})
				break
			}
		}

		// if couldn't insert, add it to the front
		if !inserted {
			l = slices.Insert(l, 0, ApplesPack{amount: apples[currentDay], expiry: expiry})
		}

		l = cleanup(l, currentDay)

		eat(l, currentDay, &eatenApplesAmount)

		currentDay++
	}

	// keep eating till no good apple packs left
	for len(l) > 0 {
		l = cleanup(l, currentDay)
		eat(l, currentDay, &eatenApplesAmount)
		currentDay++
	}

	return eatenApplesAmount
}

func main() {
	// Example 1: apples = [1,2,3,5,2], days = [3,2,1,4,2] → Expected: 7
	apples1 := []int{1, 2, 3, 5, 2}
	days1 := []int{3, 2, 1, 4, 2}
	result1 := eatenApples(apples1, days1)
	fmt.Printf("Example 1: %d (expected 7)\n", result1)

	// Example 2: apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2] → Expected: 5
	apples2 := []int{3, 0, 0, 0, 0, 2}
	days2 := []int{3, 0, 0, 0, 0, 2}
	result2 := eatenApples(apples2, days2)
	fmt.Printf("Example 2: %d (expected 5)\n", result2)

	// Example 3: apples = [2,1,10], days = [2,10,1] → Expected: 4
	apples3 := []int{2, 1, 10}
	days3 := []int{2, 10, 1}
	result3 := eatenApples(apples3, days3)
	fmt.Printf("Example 3: %d (expected 4)\n", result3)
}
