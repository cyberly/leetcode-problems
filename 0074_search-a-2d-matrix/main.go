// 74. Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	//  First pass, binary search didn't occur to me
	// Ironically works better than the recommended solution
	revPtr := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target {
			if searchArray(matrix[i], target) {
				return true
			}
		}
		if matrix[revPtr][0] <= target {
			if searchArray(matrix[revPtr], target) {
				return true
			}
		}
	}
	return false
}

func searchArray(arr []int, tar int) bool {
	revPtr := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[i] == tar || arr[revPtr] == tar {
			return true
		}
		revPtr = revPtr - 1
	}
	return false
}

func searchMatrix_binary(matrix [][]int, target int) bool {
	// Not sure if I'm an idiot but this fails due to time on LeetCode?
	rows := len(matrix) - 1
	if rows < 1 {
		return false
	}
	cols := len(matrix[0]) - 1
	l, r := 0, rows*cols-1
	for l <= r {
		mid_row := l + r/2
		mid_col := matrix[mid_row/cols][mid_row%cols]
		if target == mid_col {
			return true
		} else {
			if target < mid_col {
				r = mid_row - 1
			} else {
				l = mid_row + 1
			}
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}
