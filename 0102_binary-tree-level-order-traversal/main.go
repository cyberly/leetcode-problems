// 102. Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder_iterative(root *TreeNode) [][]int {
	levels := [][]int{}
	curr := []*TreeNode{root}
	next := []*TreeNode{}

	for {
		level := []int{}
		for _, i := range curr {
			if i == nil {
				continue
			}
			level = append(level, i.Val)
			next = append(next, i.Left, i.Right)
		}
		if len(next) == 0 {
			break
		}
		levels = append(levels, level)
		curr = next
		next = nil
	}
	return levels
}

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	if root != nil {
		recurseLevels(root, 0, &levels)
	}
	return levels
}

func recurseLevels(root *TreeNode, l int, levels *[][]int) {
	if len(*levels) == l {
		*levels = append(*levels, []int{})
	}
	(*levels)[l] = append((*levels)[l], root.Val)
	if root.Left != nil {
		recurseLevels(root.Left, l+1, levels)
	}
	if root.Right != nil {
		recurseLevels(root.Right, l+1, levels)
	}
}

func main() {
	fmt.Println("lmao")
}
