// https://leetcode.com/problems/longest-common-prefix/
package mysolution

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	for i, v := range strs[0] {
		for _, v2 := range strs[1:] {
			if i == len(v2) || v2[i] != byte(v) {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
