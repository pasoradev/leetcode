package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func recurse(root *Node, depth int) int {
	if root == nil {
		return depth
	}

	maxDepth := depth
	for _, child := range root.Children {
		maxDepth = max(recurse(child, depth+1), maxDepth)
	}
	depth = maxDepth

	return depth
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return recurse(root, 1)
}

func main() {
	node := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 2},
			{Val: 4},
		},
	}

	result := maxDepth(node)
	fmt.Println(result)
}
