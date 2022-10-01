package officaldivideandconquer

import "fmt"

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return divide(strs, 0, len(strs)-1)
}

func divide(strs []string, l, r int) string {
	fmt.Println(strs)
	if l == r {
		return strs[l]
	}
	mid := (l + r) / 2
	lcpLeft := divide(strs, l, mid)
	lcpRight := divide(strs, mid+1, r)
	return commonPrefix(lcpLeft, lcpRight)
}

func commonPrefix(left, right string) string {
	min := len(left)
	if len(left) > len(right) {
		min = len(right)
	}
	for i := 0; i < min; i++ {
		if left[i] != right[i] {
			return left[0:i]
		}
	}
	return left[0:min]
}
