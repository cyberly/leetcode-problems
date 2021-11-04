// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/

package main

import (
	"fmt"
)

func sortColors_hack(nums []int) {
	// This is terrible but 100% faster an O(2n)?
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	fmt.Println(m)
	if m[0] > 0 {
		for j := 0; j < m[0]; j++ {
			fmt.Printf("Setting nums[%v] to 0\n", j)
			nums[j] = 0
		}
	}
	if m[1] > 0 {
		for h := m[0]; h < m[0]+m[1]; h++ {
			fmt.Printf("Setting nums[%v] to 1\n", h)
			nums[h] = 1
		}
	}
	if m[2] > 0 {
		for k := m[0] + m[1]; k < len(nums); k++ {
			fmt.Printf("Setting nums[%v] to 2\n", k)
			nums[k] = 2
		}
	}
}

func sortColors(nums []int) {
	curr, p1, p2 := 0, 0, len(nums)-1
	for curr <= p2 {
		switch nums[curr] {
		case 0:
			nums[curr], nums[p1] = nums[p1], nums[curr]
			curr++
			p1++
		case 1:
			curr++
		case 2:
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	//nums := []int{2, 1}
	sortColors(nums)
	fmt.Println(nums)
}
