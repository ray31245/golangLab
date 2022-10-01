package better

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	ast := assert.New(t)
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
		},
		{
			name: "case2",
			args: args{
				nums:   []int{-9, 5, 6, -7},
				target: -16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			ast.Equal(tt.args.target, tt.args.nums[got[0]]+tt.args.nums[got[1]])
		})
	}
}
