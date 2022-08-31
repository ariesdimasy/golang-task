package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test input kasur ini rusak",
			args: args{
				input: "kasur ini rusak",
			},
			want: true,
		},
		{
			name: "test input gaga",
			args: args{
				input: "gaga",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
