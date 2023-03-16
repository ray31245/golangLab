package main

func Merge56(intervals [][]int) [][]int {
	return merge56(intervals)
}

func merge56(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	res := make([][]int, 0)
	res = append(res, intervals[0])
	curIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[curIndex][1] {
			curIndex++
			res = append(res, intervals[i])
		} else {
			res[curIndex][1] = max(intervals[i][1], res[curIndex][1])
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// func min(a int, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func partitionSort(a [][]int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if (a[j][0] < pivot[0]) || (a[j][0] == pivot[0] && a[j][1] < pivot[1]) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort(a [][]int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
