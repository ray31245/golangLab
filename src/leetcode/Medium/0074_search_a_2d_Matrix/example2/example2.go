package main

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, hight := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	for low <= hight {
		mid := low + (hight-low)>>1
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
