package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle := head
	even := false
	for head != nil {
		even = !even
		if !even {
			middle = middle.Next
		}
		head = head.Next
	}
	return middle
}

func main() {
	// Test case 1: [1,2,3,4,5]
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	result1 := middleNode(head1)
	fmt.Printf("Test 1: %d\n", result1.Val)

	// Test case 2: [1,2,3,4,5,6]
	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	result2 := middleNode(head2)
	fmt.Printf("Test 2: %d\n", result2.Val)
}
