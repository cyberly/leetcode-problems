// 451. Sort Characters By Frequency
// https://leetcode.com/problems/sort-characters-by-frequency/

package main

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	var ans string
	c := make(map[rune]int)
	b := make([][]rune, len(s)+1)
	for _, r := range s {
		c[r] += 1
	}
	for k, v := range c {
		b[v] = append(b[v], k)
	}
	for i := len(s); i > 0; i-- {
		for _, r := range b[i] {
			ans += strings.Repeat(string(r), i)
		}
	}
	return ans
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
