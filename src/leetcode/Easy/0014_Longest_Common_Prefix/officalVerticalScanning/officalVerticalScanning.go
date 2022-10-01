package officalverticalscanning

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	for i, r := range strs[0] {
		for _, refS := range strs[1:] {
			if i == len(refS) || refS[i] != byte(r) {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
