package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"server/internal/service"
	"server/internal/service/mock"
	logMock "server/pkg/logger/mock"
)

func TestValidate(t *testing.T) {
	type args struct {
		solution string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_csolution_valitated",
			args: args{
				solution: "1:5:123456:1733688600:1733692200:4bd53e4e5b3a974baf8bc7446457cf6c0f67918ad59dd86eef4b16243a7392e3:1036532",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			quoteStorage := mock.NewMockQuoteStorage(ctrl)
			tokenStorage := mock.NewMockTokenStorage(ctrl)
			l := logMock.NewMockLogger(ctrl)

			s := service.New(quoteStorage, tokenStorage, l)

			if got := s.Validate(tt.args.solution); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
