// 50, Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/

package main

import (
	"fmt"
	"sort"
)

func intersect_firstpass(nums1 []int, nums2 []int) []int {
	match := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				match = append(match, nums2[j])
				nums2[j] = 1001
				break
			}
		}
	}
	return match
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	var match []int

	for i <= len(nums1)-1 && j <= len(nums2)-1 {
		if nums1[i] == nums2[j] {
			match = append(match, nums1[i])
			i, j = i+1, j+1
			continue
		}
		if nums1[i] < nums2[j] {
			i++
			continue
		} else if nums1[i] > nums2[j] {
			j++
		}
	}
	return match
}

func main() {
	//fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
