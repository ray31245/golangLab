package mysolution

func FirstUniqChar(s string) int {
	return firstUniqChar(s)
}

type st struct {
	index  int
	repeat bool
}

func firstUniqChar(s string) int {
	m := map[rune]st{}
	for si, ss := range s {
		stIns, ok := m[ss]
		if ok {
			stIns.repeat = true
			m[ss] = stIns
		} else {
			m[ss] = st{index: si, repeat: false}
		}
	}
	// fmt.Println(len(m))
	res := len(s)
	for _, v := range m {
		// fmt.Println(string(i))
		// fmt.Println(v)
		if !v.repeat && v.index < res {
			res = v.index
		}
	}
	if res == len(s) {
		return -1
	}
	// for _, s2 := range s {
	// 	// fmt.Println(m[s2])
	// 	if !m[s2].repeat {
	// 		return m[s2].index
	// 	}
	// }
	return res
}
