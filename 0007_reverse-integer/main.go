// 7. Reverse Integer
// https://leetcode.com/problems/reverse-integer/

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ans := 0
	for x != 0 {
		p := x % 10
		x /= 10
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && p > 7) {
			return 0
		}
		if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && p < -8) {
			return 0
		}
		ans = ans*10 + p
	}
	return ans
}

func main() {
	x := 123
	fmt.Println(reverse(x))
}
