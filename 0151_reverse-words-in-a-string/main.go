// 151. Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/

package main

import (
	"fmt"
	"strings"
)

func reverseWords_slow(s string) string {
	var ans string
	buff := strings.Split(s, " ")
	for i := len(buff) - 1; i >= 0; i-- {
		if len(buff[i]) == 0 {
			continue
		}
		if len(ans) == 0 {
			ans = buff[i]
		} else {
			ans += " " + buff[i]
		}
	}
	return ans
}

func reverseWords(s string) string {
	b := strings.Fields(s)
	for i, j := 0, len(b)-1; i < j; i++ {
		b[i], b[j] = b[j], b[i]
		j--
	}
	return strings.Join(b, " ")
}

func main() {
	s := "  the  sky is blue"
	fmt.Println(reverseWords(s))
}
