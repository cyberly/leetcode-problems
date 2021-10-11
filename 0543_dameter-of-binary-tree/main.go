// 543. Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestPath(n *TreeNode, d *int) int {
	if n == nil {
		return 0
	}
	leftPath := longestPath(n.Left, d)
	rightPath := longestPath(n.Right, d)
	*d = max(*d, leftPath+rightPath)
	return max(leftPath, rightPath) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	longestPath(root, &d)
	return d
}

func main() {
	fmt.Println("bleep bloop blorp")
}
