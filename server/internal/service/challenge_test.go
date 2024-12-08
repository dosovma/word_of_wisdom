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
			name: "success_challenge_generated",
			args: args{
				r: entity.Request{
					ID:        123456,
					CreatedAt: time.Date(2024, 12, 8, 20, 10, 0, 0, time.UTC).Unix(),
				},
			},
			want: "1:5:123456:1733688600:1733692200:4bd53e4e5b3a974baf8bc7446457cf6c0f67918ad59dd86eef4b16243a7392e3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			quoteStorage := mock.NewMockQuoteStorage(ctrl)
			tokenStorage := mock.NewMockTokenStorage(ctrl)
			l := logMock.NewMockLogger(ctrl)

			s := service.New(quoteStorage, tokenStorage, l)

			if got := s.Challenge(tt.args.r); got != tt.want {
				t.Errorf("Challenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
