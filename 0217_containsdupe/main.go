// 17. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	e := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if e[nums[i]] == true {
			return true
		} else {
			e[nums[i]] = true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 4, 7, 2, 2}))
}
