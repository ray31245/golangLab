package mysolution

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	reachable := make([]bool, n)
	for _, e := range edges {
		reachable[e[1]] = true
	}
	res := []int{}
	for i, r := range reachable {
		if !r {
			res = append(res, i)
		}
	}
	return res
}
