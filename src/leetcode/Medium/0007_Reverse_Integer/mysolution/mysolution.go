package mysolution

import (
	"math"
)

func Reverse(x int) int {
	return reverse(x)
}

// https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	var reverseX int
	for x != 0 {
		reverseX = reverseX*10 + x%10
		x /= 10
	}
	if reverseX > math.MaxInt32 || reverseX < math.MinInt32 {
		return 0
	}
	return reverseX
}
