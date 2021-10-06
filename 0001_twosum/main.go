// 1. Two Sum
// https://leetcode.com/problems/two-sum/

package main

import "fmt"

// Slow but less memory
func twoSum02(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for n := i + 1; n < len(nums); n++ {
			if nums[i]+nums[n] == target && i != n {
				return []int{i, n}
			}
		}
	}
	return []int{}
}

// Single recursion but space is O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := m[diff]; ok {
			return []int{i, m[diff]}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}
