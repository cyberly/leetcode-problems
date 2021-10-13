// 232. Implement Queue using Stacks
// https://leetcode.com/problems/implement-queue-using-stacks/

package main

import "fmt"

type MyQueue struct {
	Stack []int
}

func Constructor() MyQueue {
	return MyQueue{Stack: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	pop := this.Stack[0]
	this.Stack = this.Stack[1:]
	return pop
}

func (this *MyQueue) Peek() int {
	return this.Stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.Stack) == 0
}

func main() {
	fmt.Println("placeholder")
}
