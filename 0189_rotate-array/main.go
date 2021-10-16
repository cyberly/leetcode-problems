// 189. Rotate Array
// https://leetcode.com/problems/rotate-array/

package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	l := nums[len(nums)-k:]
	r := nums[:len(nums)-k]
	copy(nums, append(l, r...))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums)
	fmt.Println(nums)
}
