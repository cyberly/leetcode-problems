// Data Structure, Day 2, Problem 2 - 88: Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1_copy := make([]int, len(nums1))
	copy(nums1_copy, nums1)
	p1 := 0
	p2 := 0
	for i := 0; i < n+m; i++ {
		if (p2 >= n) || (p1 < m && nums1_copy[p1] <= nums2[p2]) {
			nums1[i] = nums1_copy[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}
}

func merge_rev(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	for i := n + m - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
		fmt.Println(nums1)
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
