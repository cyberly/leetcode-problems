// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/

package main

import "fmt"

func moveZeroes_slow(nums []int) {
	nzIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nzIdx] = nums[i]
			nzIdx++
		}
	}
	for j := nzIdx; j < len(nums); j++ {
		nums[j] = 0
	}
}

func moveZeroes(nums []int) {
	nzIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nzIdx], nums[i] = nums[i], nums[nzIdx]
			nzIdx++
		}
	}
}

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
