package example1

func findJudge(n int, trust [][]int) int {

	// Quick response on simple case
	if n == 1 {
		return 1
	}

	// first cell is dummy, just for the convenience of indexing start from 1 to N
	trustScore := make([]int, n+1)

	// General cases:
	for _, trustNode := range trust {

		p1 := trustNode[0]
		p2 := trustNode[1]

		// decrease one point from p1 when p1 trusts other people
		trustScore[p1] -= 1

		// increase one point to p2 when p2 is trusted by others
		trustScore[p2] += 1
	}

	for personID := 1; personID <= n; personID++ {

		// town judge will be trusted by other N-1 people, and town judge trust nobody.
		if (n - 1) == trustScore[personID] {
			return personID
		}
	}

	return -1
}
