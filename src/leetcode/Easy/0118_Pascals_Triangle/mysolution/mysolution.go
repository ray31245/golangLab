package mysolution

func Generate(numRows int) [][]int {
	return generate(numRows)
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		for j := range res[i] {
			if i > 0 && j > 0 && len(res[i-1]) > j {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			} else {
				res[i][j] = 1
			}
		}
	}
	return res
}
