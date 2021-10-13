// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

package main

import "fmt"

func isValid(s string) bool {
	opens := make([]rune, 0)
	pairs := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			opens = append(opens, c)
		case ')', '}', ']':
			if len(opens) == 0 || pairs[c] != opens[len(opens)-1] {
				return false
			} else {
				opens = opens[:len(opens)-1]
			}
		}
	}
	return len(opens) == 0
}

func main() {
	fmt.Println(isValid("(){}[]"))
	fmt.Println(isValid("([)]"))

}
