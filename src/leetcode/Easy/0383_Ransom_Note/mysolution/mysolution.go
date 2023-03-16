package main

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	runeMapToCount := [26]int{}
	for _, m := range magazine {
		runeMapToCount[m-'a'] += 1
	}
	for _, r := range ransomNote {
		runeMapToCount[r-'a']--
		if runeMapToCount[r-'a'] < 0 {
			return false
		}
	}
	return true
}
