// 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertNode(root)
	return root
}

func invertNode(r *TreeNode) {
	l := r.Left
	r.Left = r.Right
	r.Right = l
	if r.Left != nil {
		invertNode(r.Left)
	}
	if r.Right != nil {
		invertNode(r.Right)
	}
}

func main() {
	fmt.Println("placeholder")
}
