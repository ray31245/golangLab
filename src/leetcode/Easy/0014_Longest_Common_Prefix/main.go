package main

import (
	// "Bevis/src/leetcode/Easy/0014_Longest_Common_Prefix/mysolution"
	"Bevis/src/leetcode/Easy/0014_Longest_Common_Prefix/officaldivideandconquer"
	"fmt"
)

func main() {
	// ans := mysolution.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	// ans := mysolution.LongestCommonPrefix([]string{"ab", "a"})
	ans := officaldivideandconquer.LongestCommonPrefix([]string{"flower", "flow", "flight", "flex", "flame"})
	fmt.Println(ans)
}
