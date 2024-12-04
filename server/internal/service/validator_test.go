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
		name      string
		args      args
		mockCalls func(args, *mock.MockQuoteStorage, *mock.MockTokenStorage)
		want      bool
	}{
		{
			name: "case_1",
			args: args{
				solution: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a:1325220",
			},
			mockCalls: func(args, *mock.MockQuoteStorage, *mock.MockTokenStorage) {

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

			tt.mockCalls(tt.args, quoteStorage, tokenStorage)

			if got := s.Validate(tt.args.solution); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
