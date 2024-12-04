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
					CreatedAt: time.Date(2024, 12, 5, 10, 10, 0, 0, time.UTC).Unix(),
				},
			},
			want: "1:5:123456:1733393400:1733397000:5c06f48d5f94ccee414ee1513a4bd75f4d89e50eac01ff45705339cff2b5148a",
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
