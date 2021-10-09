// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		if s == f {
			return true
		}
		s = s.Next
		f = f.Next.Next
	}
	return false
}

func hasCycle_hash(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	fmt.Println("lmao")
}
