// 101. Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirrored(root, root)
}

func isMirrored(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && isMirrored(l.Left, r.Right) && isMirrored(l.Right, r.Left)
}

func main() {
	fmt.Println("placeholder")
}
