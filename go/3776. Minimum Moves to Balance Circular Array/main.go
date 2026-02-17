package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func circularMod(a, m int) int {
	b := a - 1
	if b >= 0 {
		return b
	}
	return m - 1
}

func minMoves(balance []int) int64 {
	sum := 0
	negativeIndex := -1
	negativeValue := 0

	for i, val := range balance {
		if val < 0 {
			negativeIndex = i
			negativeValue = val
		} else {
			sum += val
		}
	}

	total := negativeValue + sum

	if total >= 0 && negativeIndex == -1 {
		// everything is positive
		return 0
	}
	if total < 0 {
		// sum of positive values can't make balance positive
		return -1
	}

	minMoves := int64(0)

	length := len(balance)
	leftI := circularMod(negativeIndex, length)
	rightI := (negativeIndex + 1) % length
	balanceToGo := abs(negativeValue)
	for i := range length / 2 {
		turnsItTakes := i + 1
		// check left
		if balance[leftI] > 0 {
			b := balance[leftI]
			// this person's balance is greater than needed
			if b >= balanceToGo {
				minMoves += int64(turnsItTakes * balanceToGo)
				return minMoves
			} else {
				minMoves += int64(turnsItTakes * b)
				balanceToGo -= b
			}
		}
		// check right
		if balance[rightI] > 0 {
			b := balance[rightI]
			if b >= balanceToGo {
				minMoves += int64(turnsItTakes * balanceToGo)
				return minMoves
			} else {
				minMoves += int64(turnsItTakes * b)
				balanceToGo -= b
			}
		}
		leftI = circularMod(leftI, length)
		rightI = (rightI + 1) % length
	}

	return minMoves
}

func main() {
	balance1 := []int{5, 1, -4}
	result1 := minMoves(balance1)
	fmt.Printf("Example 1: %d (expected 4)\n", result1)

	balance2 := []int{1, 2, -5, 2}
	result2 := minMoves(balance2)
	fmt.Printf("Example 2: %d (expected 6)\n", result2)

	balance3 := []int{-3, 2}
	result3 := minMoves(balance3)
	fmt.Printf("Example 3: %d (expected -1)\n", result3)

	balance4 := []int{-2, 4}
	result4 := minMoves(balance4)
	fmt.Printf("Example 4: %d (expected 2)\n", result4)

	balance5 := []int{5, 8, -7}
	result5 := minMoves(balance5)
	fmt.Printf("Example 5: %d (expected 7)\n", result5)

	balance6 := []int{14, 12, 7, -16}
	result6 := minMoves(balance6)
	fmt.Printf("Example 6: %d (expected 16)\n", result6)
}
