// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	chars := make(map[uint8]int)
	maxLen := 0
	if len(s) <= 1 {
		return len(s)
	}
	i := 0
	for j := 0; j < len(s); j++ {
		if _, ok := chars[s[j]]; ok {
			i = max(chars[s[j]], i)
		}
		maxLen = max(maxLen, j-i+1)
		chars[s[j]] = j + 1
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring(" "))        // 1
	fmt.Println(lengthOfLongestSubstring("au"))       // 2
	fmt.Println(lengthOfLongestSubstring("aab"))      // 2
}
