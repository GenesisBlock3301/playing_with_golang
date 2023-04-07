package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func insert(val int) *TreeNode {
	return &TreeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func inOrderTree(root *TreeNode) {
	if root.left != nil {
		inOrderTree(root.left)
	}
	fmt.Println(root.val)
	if root.right != nil {
		inOrderTree(root.right)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := []TreeNode{*root}
	var nextQueue []TreeNode
	for len(queue) > 0 {
		var level []int
		for _, node := range queue {
			level = append(level, node.val)
			if node.left != nil {
				nextQueue = append(nextQueue, *node.left)
			}
			if node.right != nil {
				nextQueue = append(nextQueue, *node.right)
			}
		}
		result = append(result, level)
		queue = nextQueue
		nextQueue = []TreeNode{}
	}
	return result
}
func main() {
	//arr := []int{3, 9, 20, 0, 0, 15, 7}
	root := insert(3)
	root.left = insert(9)
	root.right = insert(20)
	root.right.left = insert(15)
	root.right.right = insert(7)
	//inOrderTree(root)

	result := levelOrder(root)
	println(result)
}
