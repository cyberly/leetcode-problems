// ind-duplicate-file-in-system/
// https://leetcode.com/problems/find-duplicate-file-in-system/

package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	bucket := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		folder := parts[0]
		for i := 1; i < len(parts); i++ {
			cursor := strings.IndexByte(parts[i], '(')
			bucket[parts[i][cursor:]] = append(bucket[parts[i][cursor:]], folder+"/"+parts[i][:cursor])
		}
	}
	ans := make([][]string, 0, len(bucket))
	for _, files := range bucket {
		if len(files) >= 2 {
			ans = append(ans, files)
		}
	}
	return ans
}

func main() {
	paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}
	fmt.Println(findDuplicate(paths))
}
