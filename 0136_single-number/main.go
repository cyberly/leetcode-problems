// 136. Single Number
// https://leetcode.com/problems/single-number/

package main

import "fmt"

func singleNumber_map(nums []int) int {
	val := 0
	counts := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := counts[nums[i]]; ok {
			delete(counts, nums[i])
		} else {
			counts[nums[i]] = 1
		}
	}
	for k := range counts {
		val = k
	}
	return val
}

func singleNumber(nums []int) int {
	// Use bitwise, single iteration
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		nums[0] = nums[0] ^ nums[i]
		fmt.Println(nums[0])
	}
	return nums[0]
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
