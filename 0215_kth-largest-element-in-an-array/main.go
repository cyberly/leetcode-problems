// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/

package main

import (
	"fmt"
	"sort"
)

func findKthLargest_sort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func qSelect(nums []int, left, right, k int) int {
	// I don't actually understand this
	// And it is slower than the sort above?
	// Whatever.
	if left == right {
		return nums[left]
	}
	p := right
	l := left
	r := right - 1
	for l <= r {
		for l <= r && nums[l] >= nums[p] {
			l++
		}
		for l <= r && nums[r] < nums[p] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[l], nums[p] = nums[p], nums[l]
	p = l
	if p > k-1 {
		return qSelect(nums, left, p-1, k)
	} else {
		if p < k-1 {
			return qSelect(nums, p+1, right, k)
		}
	}
	return nums[p]
}

func findKthLargest(nums []int, k int) int {
	return qSelect(nums, 0, len(nums)-1, k)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
