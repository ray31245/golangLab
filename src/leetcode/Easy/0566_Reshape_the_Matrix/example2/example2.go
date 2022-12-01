package main

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	return matrixReshape(mat, r, c)
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if canReshape(mat, r, c) {
		return reshape(mat, r, c)
	}
	return mat
}

func canReshape(nums [][]int, r, c int) bool {
	row := len(nums)
	colume := len(nums[0])
	return row*colume == r*c
}

func reshape(nums [][]int, r, c int) [][]int {
	newShape := make([][]int, r)
	for index := range newShape {
		newShape[index] = make([]int, c)
	}
	rowIndex, colIndex := 0, 0
	for _, row := range nums {
		for _, col := range row {
			if colIndex == c {
				colIndex = 0
				rowIndex++
			}
			newShape[rowIndex][colIndex] = col
			colIndex++
		}
	}
	return newShape
}
