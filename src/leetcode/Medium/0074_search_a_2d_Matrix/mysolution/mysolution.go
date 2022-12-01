package mysolution

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	l := len(matrix[0])
	start := 0
	end := len(matrix)*l - 1
	var mid int
	for start <= end {
		mid = end / 2
		if matrix[start/l][start%l] == target || matrix[end/l][end%l] == target {
			return true
		}
		// time.Sleep(time.Second)
		// fmt.Println("end: ", end)
		// fmt.Println(matrix[end/l][end%l])
		// fmt.Println("start: ", start)
		// fmt.Println(matrix[start/l][start%l])
		// fmt.Println("--------")
		if start == end {
			break
		}
		if target > matrix[mid/l][mid%l] {
			start = mid
			end -= 1
		} else {
			end = mid
			start += 1
		}
	}
	return false
}
