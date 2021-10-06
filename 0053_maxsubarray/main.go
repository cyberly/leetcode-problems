//53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
//
// TODO: Slow result, figure out why

package main

import (
	"fmt"
)

func maxSubArraySlow(nums []int) int {
	subSum := nums[0]
	currentSum := 0

	for _, i := range nums {
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum = currentSum + i
		if currentSum > subSum {
			subSum = currentSum
		}
	}
	return subSum
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		currentSum = Max(currentSum+nums[i], nums[i])
		maxSum = Max(currentSum, maxSum)
	}
	return maxSum
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
