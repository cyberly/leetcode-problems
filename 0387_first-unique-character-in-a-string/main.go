// 387. First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/

package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	chars := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}
	for j := 0; j < len(s); j++ {
		if chars[s[j]] == 1 {
			return j
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
