// 701. Insert into a Binary Search Tree
// https://leetcode.com/problems/insert-into-a-binary-search-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST_iteration(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	curr := root
	newNode := &TreeNode{
		Val: val,
	}
	for curr != nil {
		if curr.Val > val {
			if curr.Left != nil {
				curr = curr.Left
				continue
			} else {
				curr.Left = newNode
				curr = nil
			}
		} else {
			if curr.Right != nil {
				curr = curr.Right
				continue
			} else {
				curr.Right = newNode
				curr = nil
			}
		}
	}

	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func main() {
	fmt.Println("placeholder")
}
