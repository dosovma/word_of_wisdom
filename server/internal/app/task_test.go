package app

import (
	"testing"
	"time"
)

func Test_taskGenerator(t *testing.T) {
	type args struct {
		requestID   int64
		requestTime int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case_1",
			args: args{
				requestID:   123456,
				requestTime: time.Date(2024, 12, 1, 10, 10, 0, 0, time.UTC).Unix(),
			},
			want: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task(tt.args.requestID, tt.args.requestTime); got != tt.want {
				t.Errorf("taskGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
