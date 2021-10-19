// 344. Reverse String
// https://leetcode.com/problems/reverse-string/

package main

import "fmt"

func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}
