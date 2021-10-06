package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	newMat := make([][]int, r)
	for i := range newMat {
		newMat[i] = make([]int, c)
	}
	rPtr, cPtr := 0, 0
	for _, row := range mat {
		for _, col := range row {
			if cPtr == c {
				cPtr = 0
				rPtr++
			}
			newMat[rPtr][cPtr] = col
			cPtr++
		}
	}
	return newMat
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2, 5, 6}, {3, 4, 7, 8}}, 4, 2))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}
