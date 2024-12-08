package solver_test

/*
DO NOT USE THIS TEST
IT'S TO MANUALLY GENERATE HEADER

func TestSolv_Solve(t *testing.T) {
	type args struct {
		challenge string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				challenge: "1:5:123456:1733688600:1733692200:4bd53e4e5b3a974baf8bc7446457cf6c0f67918ad59dd86eef4b16243a7392e3",
			},
			want:    "1:5:123456:1733688600:1733692200:4bd53e4e5b3a974baf8bc7446457cf6c0f67918ad59dd86eef4b16243a7392e3:1036532",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			log := loggerMock.NewMockLogger(ctrl)

			so := &Solv{
				log: log,
			}
			got, err := so.Solve(tt.args.challenge)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
