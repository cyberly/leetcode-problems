// 653. Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	return search(root, m, k)
}

func search(node *TreeNode, m map[int]bool, k int) bool {
	if node == nil {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	} else {
		m[node.Val] = true
	}
	return search(node.Left, m, k) || search(node.Right, m, k)
}

func test(m map[int]bool) {
	m[10] = true
}

func main() {
	fmt.Println("placeholder")
}
