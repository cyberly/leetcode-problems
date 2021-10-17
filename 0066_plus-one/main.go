// 66. Plus One
// https://leetcode.com/problems/plus-one/

package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i] += 1
		return digits
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}
