// 203. Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/

package main

import "fmt"

var start *ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sent := &ListNode{
		Val:  0,
		Next: head,
	}
	prev, curr := sent, head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return sent.Next
}

func main() {
	fmt.Println("placeholder")
}
