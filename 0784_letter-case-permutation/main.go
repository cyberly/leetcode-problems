// 784. Letter Case Permutation
// https://leetcode.com/problems/letter-case-permutation/

package main

import (
	"fmt"
	"unicode"
)

func letterCasePermutation_slow(s string) []string {
	ans := []string{s}
	for i := 0; i < len(s); i++ {
		curr := len(ans)
		for j := 0; j < curr; j++ {
			str := []rune(ans[j])
			if unicode.IsLetter(str[i]) {
				if unicode.IsUpper(str[i]) {
					str[i] = unicode.ToLower(str[i])
				} else {
					str[i] = unicode.ToUpper(str[i])
				}
				ans = append(ans, string(str))
			}
		}
	}
	return ans
}

func letterCasePermutation(s string) []string {
	// Can be quite a bit faster when being evaluated
	// Cleaner loop, drop unicode
	ans := []string{s}
	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			len := len(ans)
			for j := 0; j < len; j++ {
				p := []rune(ans[j])
				p[i] = r + 'A' - 'a'
				ans = append(ans, string(p))
			}
		}
		if r >= 'A' && r <= 'Z' {
			len := len(ans)
			for j := 0; j < len; j++ {
				p := []rune(ans[j])
				p[i] = r + 'a' - 'A'
				ans = append(ans, string(p))
			}
		}
	}
	return ans
}

func main() {
	s := "a1b2"
	//s = "C"
	fmt.Println(letterCasePermutation(s))
}
