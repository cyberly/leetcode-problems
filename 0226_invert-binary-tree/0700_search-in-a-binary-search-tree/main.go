// 700. Search in a Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST_iteration(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Left != nil && root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func main() {
	fmt.Println("placeholder")
}
