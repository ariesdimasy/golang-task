package main

import (
	"testing"
)

func TestVocalKonsonan(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test input kasur ini rusak",
			args: args{
				input: "a",
			},
			want: "vocal",
		},
		{
			name: "test input gaga",
			args: args{
				input: "b",
			},
			want: "konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VocalKonsonan(tt.args.input); got != tt.want {
				t.Errorf("VocalKonsonan() = %v, want %v", got, tt.want)
			}
		})
	}
}
