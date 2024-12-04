package service

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		task string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case_1",
			args: args{
				task: "1:5:79499:1733323676:1733410076:5a0dd71db927c6f90f23417b9f8e0668250ccdca525aa6101815101491cc4f97",
			},
			want: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a:1325220",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve() got = %v, want %v", got, tt.want)
			}
		})
	}
}
