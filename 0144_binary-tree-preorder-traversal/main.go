// 144. Binary Tree Preorder Traversal
// https://leetcode.com/problems/binary-tree-preorder-traversal/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal_iterative(root *TreeNode) []int {
	stack := []*TreeNode{root}
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return ans
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal(root.Left)...)
	ans = append(ans, preorderTraversal(root.Right)...)
	return ans
}

func main() {
	fmt.Println("place")
}
