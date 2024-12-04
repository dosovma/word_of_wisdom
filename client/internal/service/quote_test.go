package service_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"client/internal/service"
	"client/internal/service/mock"
)

func TestService_Quote(t *testing.T) {
	type args struct {
		requestID   int64
		requestTime int64
	}

	tests := []struct {
		name      string
		args      args
		mockCalls func(args, *mock.MockTCPClient, *mock.MockSolver)
		want      string
		wantErr   bool
	}{
		{
			name: "success_get_quote",
			args: args{
				requestID:   12345,
				requestTime: 1734269872,
			},
			mockCalls: func(a args, c *mock.MockTCPClient, s *mock.MockSolver) {
				reqID := strconv.FormatInt(a.requestID, 10)
				reqTime := strconv.FormatInt(a.requestTime, 10)

				challenge := "1:5:123456:1733393400:1733397000:5c06f48d5f94ccee414ee1513a4bd75f4d89e50eac01ff45705339cff2b5148a"
				solution := "1:5:123456:1733393400:1733397000:5c06f48d5f94ccee414ee1513a4bd75f4d89e50eac01ff45705339cff2b5148a:987222"
				token := uuid.New().String()

				gomock.InOrder(
					c.EXPECT().GetChallenge(reqID, reqTime).Return(challenge, nil).Times(1),
					s.EXPECT().Solve(challenge).Return(solution, nil).Times(1),
					c.EXPECT().GetTokenBySolution(solution).Return(token, nil).Times(1),
					c.EXPECT().GetQuote(token).Return("quote", nil).Times(1),
				)
			},
			want:    "quote",
			wantErr: false,
		},
		{
			name: "error_get_token_by_solution",
			args: args{
				requestID:   12345,
				requestTime: 1734269872,
			},
			mockCalls: func(a args, c *mock.MockTCPClient, s *mock.MockSolver) {
				reqID := strconv.FormatInt(a.requestID, 10)
				reqTime := strconv.FormatInt(a.requestTime, 10)

				challenge := "1:5:123456:1733393400:1733397000:5c06f48d5f94ccee414ee1513a4bd75f4d89e50eac01ff45705339cff2b5148a"
				solution := "1:5:123456:1733393400:1733397000:5c06f48d5f94ccee414ee1513a4bd75f4d89e50eac01ff45705339cff2b5148a:987222"

				gomock.InOrder(
					c.EXPECT().GetChallenge(reqID, reqTime).Return(challenge, nil).Times(1),
					s.EXPECT().Solve(challenge).Return(solution, nil).Times(1),
					c.EXPECT().GetTokenBySolution(solution).Return("", errors.New("error")).Times(1),
				)
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tcpClient := mock.NewMockTCPClient(ctrl)
			solver := mock.NewMockSolver(ctrl)

			s := service.NewService(tcpClient, solver)

			tt.mockCalls(tt.args, tcpClient, solver)

			got, err := s.Quote(tt.args.requestID, tt.args.requestTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("Quote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Quote() got = %v, want %v", got, tt.want)
			}
		})
	}
}
