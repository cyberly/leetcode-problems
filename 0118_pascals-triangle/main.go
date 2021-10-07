// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/

package main

import "fmt"

func generate(numRows int) [][]int {
	tri := make([][]int, numRows)
	for r := 0; r < len(tri); r++ {
		tri[r] = make([]int, r+1)
		tri[r][0], tri[r][r] = 1, 1
		val := 1
		for v := 1; v < r; v++ {
			val = val * (r - v + 1) / v
			tri[r][v] = val
		}
	}
	return tri
}

func main() {
	fmt.Println(generate(5))
}
