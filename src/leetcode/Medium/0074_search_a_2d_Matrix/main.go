package main

import (
	"Bevis/src/leetcode/Medium/0074_search_a_2d_Matrix/mysolution"
	"fmt"
)

func main() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(mysolution.SearchMatrix(matrix, 7))
}
