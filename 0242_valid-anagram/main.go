// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

package main

import (
	"fmt"
	"sort"
)

// Fastest
func isAnagram_hash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}
	for j := 0; j < len(t); j++ {
		if chars[t[j]] < 1 {
			return false
		}
		chars[t[j]]--
	}
	return true
}

// Unnecessary but ok
// Impletment necessary methods for sort.Sort
// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
// https://pkg.go.dev/sort#Sort
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// While kind of cute I guess, this method sort of sucks
// Might be better to split the string into a sorted array but
// map still seems faster in Go idk
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s = sortString(s)
	t = sortString(t)
	return s == t
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
