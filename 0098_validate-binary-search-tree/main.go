// 98. Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return recurse(root, nil, nil)
}

func recurse(root, l, r *TreeNode) bool {
	if root == nil {
		return true
	}
	lValid, rValid := true, true
	if l != nil {
		lValid = root.Val > l.Val
	}
	if r != nil {
		rValid = root.Val < r.Val
	}
	return lValid && rValid && recurse(root.Left, l, root) && recurse(root.Right, root, r)
}

func main() {
	fmt.Println("placeholder")
}
