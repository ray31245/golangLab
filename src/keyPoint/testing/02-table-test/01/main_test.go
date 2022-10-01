package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		name   string
		data   []int
		answer int
	}
	tests := []test{
		{"1", []int{21, 21}, 42},
		{"2", []int{3, 4, 5}, 12},
		{"3", []int{1, 1}, 2},
		{"4", []int{-1, 0, 1}, 0},
		// {"5", []int{-1, 0, 1}, 7},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("EXpected", v.answer, "Got", x)
		}
	}
}

// func Test_mySum(t *testing.T) {
// 	type args struct {
// 		xi []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want interface{}
// 	}{
// 		{
// 			name: "aaaa",
// 			args: args{
// 				xi: []int{1, 2, 3},
// 			},
// 			want: 6,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := mySum(tt.args.xi...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("mySum() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
