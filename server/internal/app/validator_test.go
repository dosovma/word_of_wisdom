package app

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		response string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case_1",
			args: args{
				response: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a:1325220",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.response); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
