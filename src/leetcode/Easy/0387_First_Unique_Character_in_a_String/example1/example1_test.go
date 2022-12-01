package example1

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	n string
}

type ans struct {
	one int
}

func Test_firstUniqChar(t *testing.T) {
	qs := []question{
		{
			para: para{"leetcode"},
			ans:  ans{0},
		},
		{
			para: para{"loveleetcode"},
			ans:  ans{2},
		},
		{
			para: para{"aaa"},
			ans:  ans{-1},
		},
	}

	fmt.Println("---------Leetcode problem 387------------------")
	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("input: %v output: %v \n", p, firstUniqChar(p.n))
	}
	fmt.Printf("\n\n\n")
}
