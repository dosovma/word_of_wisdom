package service_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"server/internal/service"
	"server/internal/service/entity"
	"server/internal/service/mock"
	logMock "server/pkg/logger/mock"
)

func TestService_Challenge(t *testing.T) {
	type args struct {
		r entity.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case_1",
			args: args{
				r: entity.Request{
					ID:        123456,
					CreatedAt: time.Date(2024, 12, 1, 10, 10, 0, 0, time.UTC).Unix(),
				},
			},
			want: "1:5:123456:1733047800:1733134200:387a824155f23c7392190f2ac624a43270ac6bde6b21384da1ee1dcdf2644b9a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			quoteStorage := mock.NewMockQuoteStorage(ctrl)
			tokenStorage := mock.NewMockTokenStorage(ctrl)
			l := logMock.NewMockLogger(ctrl)

			s := service.New(quoteStorage, tokenStorage, l)

			//tt.mockCalls(tt.args, quoteStorage, tokenStorage)

			if got := s.Challenge(tt.args.r); got != tt.want {
				t.Errorf("Challenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
