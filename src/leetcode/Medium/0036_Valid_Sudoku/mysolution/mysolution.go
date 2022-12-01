package mysolution

func IisValidSudoku(board [][]byte) bool {
	return isValidSudoku(board)
}

type numExist map[int]map[byte]bool

func isValidSudoku(board [][]byte) bool {
	row := initMap()
	col := initMap()
	block := initMap()
	var b int
	for i, v := range board {
		for j, x := range v {
			b = choiceBlock(i, j)
			if row[i][x] || col[j][x] || block[b][x] {
				return false
			}
			if x != byte('.') {
				row[i][x] = true
				col[j][x] = true
				block[b][x] = true
			}
		}
	}
	return true
}

func choiceBlock(x, y int) int {
	bX := x / 3
	bY := y / 3
	return bY + bX*3
}

func initMap() numExist {
	m := make(numExist, 9)
	for i := 0; i < 9; i++ {
		m[i] = make(map[byte]bool, 9)
		for j := 1; j < 9; j++ {
			m[i][byte(j)] = false
		}
	}
	return m
}
