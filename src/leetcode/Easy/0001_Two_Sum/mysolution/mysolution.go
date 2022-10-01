package mysolution

import "fmt"

func twoSum(nums []int, target int) []int {
	match := []int{}
	for nk, n := range nums {
		for mk, m := range match {
			if n == m {
				return []int{nk, mk}
			}
		}
		match = append(match, target-n)
		fmt.Println("match=", match)
	}
	return []int{}
}

// O(n^2/2)
// O(n)
