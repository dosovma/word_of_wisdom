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
				solution: "1:5:123456:1733393400:1733397000:5c06f48d5f94ccee414ee1513a4bd75f4d89e50eac01ff45705339cff2b5148a:987222",
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
