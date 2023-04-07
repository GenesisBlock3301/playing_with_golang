package main

import "fmt"

type RNode struct {
	Value int
	Left  *RNode
	Right *RNode
}

func findLargestValue(root *RNode) int {
	if root == nil {
		return 0
	}
	findLeftLargest := findLargestValue(root.Left)
	findRightLargest := findLargestValue(root.Right)
	if root.Value > findLeftLargest && root.Value > findRightLargest {
		return root.Value
	} else if findLeftLargest > findRightLargest {
		return findLeftLargest
	} else {
		return findRightLargest
	}
}

func main() {
	root := &RNode{
		Value: 5,
		Left: &RNode{
			Value: 3,
			Left: &RNode{
				Value: 1,
			},
			Right: &RNode{
				Value: 4,
			},
		},
		Right: &RNode{
			Value: 7,
			Right: &RNode{
				Value: 9,
			},
		},
	}
	largest := findLargestValue(root)
	fmt.Println(largest)
}
