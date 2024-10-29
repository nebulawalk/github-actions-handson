package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "偶数の場合",
			args: args{number: 2},
			want: "even",
		},
		{
			name: "奇数の場合",
			args: args{number: 3},
			want: "odd",
		},
		{
			name: "0の場合",
			args: args{number: 0},
			want: "even",
		},
		{
			name: "負の偶数の場合",
			args: args{number: -2},
			want: "even",
		},
		{
			name: "負の奇数の場合",
			args: args{number: -3},
			want: "odd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenOrOdd(tt.args.number); got != tt.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
