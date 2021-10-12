// 374. Guess Number Higher or Lower
// https://leetcode.com/problems/guess-number-higher-or-lower/

package main

import (
	"fmt"
	"math"
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(n int) int {
	return n
}

func guessNumber_binary(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == -1 {
			right = mid - 1
		}
		if res == 1 {
			left = mid + 1
		}
		if res == 0 {
			return mid
		}
	}
	return -1
}

func guessNumber(n int) int {
	left := 1
	right := n
	for left <= right {
		m1 := left + (right-left)/3
		m2 := left + 2*((right-left)/3)
		r1 := guess(m1)
		r2 := guess(m2)
		// Long but I like the clarity
		if r1 == 0 {
			return m1
		}
		if r2 == 0 {
			return m2
		}
		if r1 == -1 {
			right = m1 - 1
		}
		if r1 == 1 {
			if r2 == 1 {
				left = m2 + 1
			}
			if r2 == -1 {
				left = m1 + 1
				right = m2 - 1
			}
		}
	}
	return -1
}

func main() {
	top := math.MaxInt32
	fmt.Println(int(top))
	fmt.Println("lmao")
}
