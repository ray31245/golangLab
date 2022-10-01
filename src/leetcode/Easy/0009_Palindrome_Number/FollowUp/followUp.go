package FollowUp

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var reverseX int
	refX := x
	for refX != 0 {
		reverseX = reverseX*10 + refX%10
		refX /= 10
	}
	return reverseX == x
}
