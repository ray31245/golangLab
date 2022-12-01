package buildinsearch

import (
	"sort"
)

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target < matrix[0][0] || matrix[m-1][n-1] < target {
		return false
	}

	// 尋找行
	r := 0
	for r < m && matrix[r][0] <= target {
		r++
	}
	r--

	// 二分搜尋法 target
	res := sort.Search(len(matrix[r]), func(i int) bool {
		return matrix[r][i] >= target
	})
	return res < len(matrix[r]) && matrix[r][res] == target
}
