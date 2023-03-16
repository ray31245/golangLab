package main

import (
	// "Bevis/src/leetcode/Easy/0014_Longest_Common_Prefix/mysolution"
	"Bevis/src/leetcode/Easy/0014_Longest_Common_Prefix/mysolution"
	"fmt"
)

func main() {
	// ans := mysolution.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	// ans := mysolution.LongestCommonPrefix([]string{"ab", "a"})
	ans := mysolution.LongestCommonPrefix([]string{"flower", "flow", "flight", "flex", "flame"})
	fmt.Println(ans)
}
