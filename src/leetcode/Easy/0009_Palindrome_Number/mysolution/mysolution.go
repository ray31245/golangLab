package mysolution

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	length := len(strX) - 1
	for i := 0; i < length; i++ {
		if strX[i] != strX[length-i] {
			return false
		}
	}
	return true
}

// func isPalindrome(x int) bool {
// 	strX := strconv.Itoa(x)
//     length := (len(strX) - 1)>>1
// 	for i := 0; i < length; i++ {
// 		if strX[i] != strX[length-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
