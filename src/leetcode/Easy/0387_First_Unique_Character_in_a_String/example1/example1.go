package example1

func FirstUniqChar(s string) int {
	return firstUniqChar(s)
}

func firstUniqChar(s string) int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
