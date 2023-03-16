package main

func Merge(intervals [][]int) [][]int {
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	for _, v := range intervals {
		if len(res) == 0 {
			res = append(res, v)
			continue
		}
		lastRes := res[len(res)-1]
		if lastRes[1] >= v[0] {
			if v[1] > lastRes[1] {
				lastRes[1] = v[1]
			}
			if v[0] < lastRes[0] {
				lastRes[0] = v[0]
			}
		} else {
			res = append(res, v)
		}
	}

	return res
}
