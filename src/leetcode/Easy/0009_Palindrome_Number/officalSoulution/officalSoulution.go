package officalsoulution

import "fmt"

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	// 如果x為零，必定不會回文。
	// 另外如果最後一位數為零，第一位數也必須為零，然而規則上第一位數為零會被去掉
	// 承上條規則的例外狀況，如果x為零是為回文
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var reverseX int
	for reverseX < x {
		reverseX = reverseX*10 + x%10
		x = x / 10
	}
	// 分為兩種狀況: 奇數或是偶數
	// 奇數: 去掉中間的數字 (reverseX/10)，例如輸入為12321最後的結果會得到x = 12, reverseX = 123,
	// 偶數: 直接判斷反轉前後的各一半的數字是否相同
	fmt.Println(reverseX)
	fmt.Println(x)
	return reverseX == x || x == reverseX/10
}
