// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome2(s string) bool {
	rgx, _ := regexp.Compile("[^A-Za-z0-9]+")
	t := strings.ToLower(rgx.ReplaceAllString(s, ""))
	j := len(t) - 1
	for i := 0; i < j; i++ {
		if t[i] != t[j] {
			return false
		}
		j--
	}
	return true
}

func isPalindrome(s string) bool {
	// 100% cause regex is slow
	b := []rune(s)
	for l, r := 0, len(s)-1; l <= r; {
		if !unicode.IsLetter(b[l]) && !unicode.IsNumber(b[l]) {
			l++
			continue
		} else {
			b[l] = unicode.ToLower(b[l])
		}
		if !unicode.IsLetter(b[r]) && !unicode.IsNumber(b[r]) {
			r--
			continue
		} else {
			b[r] = unicode.ToLower(b[r])
		}
		if b[l] != b[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
