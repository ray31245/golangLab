package mysolution

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "121",
			args: args{121},
			want: true,
		},
		{
			name: "121121",
			args: args{121121},
			want: true,
		},
		{
			name: "123",
			args: args{123},
			want: false,
		},
		{
			name: "-121",
			args: args{-121},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
