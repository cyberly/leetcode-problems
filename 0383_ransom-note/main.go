// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/

package main

import (
	"fmt"
)

func canConstruct_2hash(ransomNote string, magazine string) bool {
	rChars := make(map[uint8]int)
	mChars := make(map[uint8]int)
	if len(ransomNote) > len(magazine) {
		return false
	}
	for i := 0; i < len(magazine); i++ {
		if i < len(ransomNote) {
			rChars[ransomNote[i]]++
		}
		mChars[magazine[i]]++
	}
	for k := range rChars {
		if rChars[k] > mChars[k] {
			return false
		}
	}
	return true
}

func canConstruct(ransomNote string, magazine string) bool {
	// Not a huge difference but a bit cleaner I guess
	mChars := make(map[uint8]int)
	if len(ransomNote) > len(magazine) {
		return false
	}
	for i := 0; i < len(magazine); i++ {
		mChars[magazine[i]]++
	}
	for j := 0; j < len(ransomNote); j++ {
		if mChars[ransomNote[j]] < 1 {
			return false
		}
		mChars[ransomNote[j]]--
	}
	return true
}

func main() {
	fmt.Println(canConstruct("fffbfg", "effjfggbffjdgbjjhhdegh"))
}
