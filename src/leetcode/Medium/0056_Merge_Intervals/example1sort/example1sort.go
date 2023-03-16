package main

import "sort"

func Merge56(intervals [][]int) [][]int {
	return merge56(intervals)
}

func merge56(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
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
