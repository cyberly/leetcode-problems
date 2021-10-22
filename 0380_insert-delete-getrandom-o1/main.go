// 380. Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/

package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	n map[int]int
	k []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.n[val]; !ok {
		this.k = append(this.k, val)
		this.n[val] = len(this.k) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	v, ok := this.n[val]
	if !ok {
		return false
	}
	delete(this.n, val)
	if v != len(this.k)-1 {
		this.n[this.k[len(this.k)-1]] = v
		this.k[v] = this.k[len(this.k)-1]
	}
	this.k = this.k[:len(this.k)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Int() % len(this.k)
	return this.k[r]
}

func main() {
	nums := Constructor()
	fmt.Println(nums.Insert(3))
	fmt.Println(nums.Remove(2))
	fmt.Println(nums.GetRandom())
}
