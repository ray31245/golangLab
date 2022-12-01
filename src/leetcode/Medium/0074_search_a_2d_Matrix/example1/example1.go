package main

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
	i, j := 0, n-1
	for i <= j {
		mid := (i + j) >> 1
		switch {
		case matrix[r][mid] < target:
			i = mid + 1
		case target < matrix[r][mid]:
			j = mid - 1
		default:
			return true
		}
	}
	return matrix[r][j] == target
}
