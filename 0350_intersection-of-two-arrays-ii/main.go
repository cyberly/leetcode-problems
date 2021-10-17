// 50, Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/

package main

import (
	"fmt"
	"sort"
)

func intersect_sort(nums1 []int, nums2 []int) []int {
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

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	countMap := make(map[int]int)
	ans := []int{}
	for i := 0; i < len(nums1); i++ {
		countMap[nums1[i]] += 1
	}
	for j := 0; j < len(nums2); j++ {
		if c, ok := countMap[nums2[j]]; ok && c > 0 {
			ans = append(ans, nums2[j])
			countMap[nums2[j]] -= 1
		}
	}
	return ans
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
