// 145. Binary Tree Postorder Traversal
// https://leetcode.com/problems/binary-tree-postorder-traversal/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal_iterative(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ans := []int{}
	if root == nil {
		return ans
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		root := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append([]int{root.Val}, ans...)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	return ans
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	left = append(left, right...)
	left = append(left, root.Val)

	return left
}

func main() {
	fmt.Println("lmao")
}
