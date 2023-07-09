package main

func IsAnagram(s string, t string) bool {
	return isAnagram(s, s)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cnt [26]int
	for _, v := range s {
		cnt[v-'a']++
	}
	for _, v := range t {
		cnt[v-'a']--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}
