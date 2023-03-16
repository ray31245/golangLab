package main

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomTally := make([]int, 26)
	for _, ch := range ransomNote {
		ransomTally[ch-'a']++
	}

	for _, ch := range magazine {
		ransomTally[ch-'a']--
	}
	for _, t := range ransomTally {
		if t > 0 {
			return false
		}
	}
	return true
}
