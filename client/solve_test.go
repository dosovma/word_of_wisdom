package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		task string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case_1",
			args: args{
				task: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a",
			},
			want: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a:1325220",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.task); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
