// 36. Valid Sudoku
// https://leetcode'.'com/problems/valid-sudoku/

package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku_on3(board [][]byte) bool {
	// Start with rpws
	for r := 0; r < 9; r++ {
		rExists := [10]int{}
		for c := 0; c < 9; c++ {
			cellVal := board[r][c]
			if string(cellVal) != "." {
				value, _ := strconv.Atoi(string(cellVal))
				if value > 9 || value < 1 {
					return false
				}
				if rExists[value] == 1 {
					return false
				}
				rExists[value] = 1
			}
		}

	}

	// Columns next
	for r := 0; r < 9; r++ {
		cExists := [10]int{}
		for c := 0; c < 9; c++ {
			cellVal := board[c][r]
			if string(cellVal) != "." {
				value, _ := strconv.Atoi(string(cellVal))
				if value > 9 || value < 1 {
					return false
				}
				if cExists[value] == 1 {
					return false
				}
				cExists[value] = 1
			}
		}
	}

	//box
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			bExists := [10]int{}
			for rr := r * 3; rr < r*3+3; rr++ {
				for cc := c * 3; cc < c*3+3; cc++ {
					cellVal := board[rr][cc]
					if string(cellVal) != "." {
						value, _ := strconv.Atoi(string(cellVal))
						if value > 9 || value < 1 {
							return false
						}
						if bExists[value] == 1 {
							return false
						}
						bExists[value] = 1
					}
				}
			}

		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	rBuff, cBuff, bBuff := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rBuff[i] = make([]bool, 10)
		cBuff[i] = make([]bool, 10)
		bBuff[i] = make([]bool, 10)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cellVal := board[r][c]
			if string(cellVal) != "." {
				value, _ := strconv.Atoi(string(cellVal))
				if rBuff[r][value] || cBuff[c][value] || bBuff[r/3*3+c/3][value] {
					return false
				}
				rBuff[r][value] = true
				cBuff[c][value] = true
				bBuff[r/3*3+c/3][value] = true
			}
		}
	}
	return true
}

func main() {
	/* fmt.Println(isValidSudoku([][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '6', '.'},
	{'6', '8', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '.', '.', '.', '.', '.', '.', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}})) */
	fmt.Println(isValidSudoku([][]byte{
		{'5', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}}))
}
