// 130. Surrounded Regions
// https://leetcode.com/problems/surrounded-regions/

package main

import "fmt"

func DFS(board [][]byte, r, c int) {
	if board[r][c] != 'O' {
		return
	}
	board[r][c] = 'S'
	// Check N, S, E, W
	if r > 0 {
		DFS(board, r-1, c)
	}
	if r < len(board)-1 {
		DFS(board, r+1, c)
	}
	if c < len(board[0])-1 {
		DFS(board, r, c+1)
	}
	if c > 0 {
		DFS(board, r, c-1)
	}
}

func solve(board [][]byte) {
	if len(board) <= 1 || len(board[0]) <= 1 {
		return
	}
	// Check borders - T, B, L, R
	for i := 0; i < len(board[0]); i++ {
		DFS(board, 0, i)
	}
	for i := 0; i < len(board[0]); i++ {
		DFS(board, len(board)-1, i)
	}
	for i := 1; i < len(board)-1; i++ {
		DFS(board, i, 0)
	}
	for i := 1; i < len(board)-1; i++ {
		DFS(board, i, len(board[0])-1)
	}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
			if board[r][c] == 'S' {
				board[r][c] = 'O'
			}
		}
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'x'}}
	solve(board)
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
}
