package client_test

import (
	"net"
	"testing"

	"github.com/golang/mock/gomock"

	"client/internal/infrastructure/client"
	"client/internal/infrastructure/client/mock"
	loggerMock "client/pkg/logger/mock"
)

func TestClient_GetChallenge(t *testing.T) {
	type args struct {
		requestID   string
		requestTime string
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(net.Conn, *mock.MockMessenger, *loggerMock.MockLogger)
		want      string
		wantErr   bool
	}{
		{
			name: "success_get_challenge",
			args: args{
				requestID:   "123456",
				requestTime: "1732456988",
			},
			mockCalls: func(conn net.Conn, messenger *mock.MockMessenger, logger *loggerMock.MockLogger) {
				data := []string{"X-Command:Token", "X-Request-id:123456", "X-Request-time:1732456988"}

				gomock.InOrder(
					messenger.EXPECT().Write(conn, data).Return(nil).Times(1),
					logger.EXPECT().Println("token requested"),
					messenger.EXPECT().Read(conn).Return([]string{"X-Challenge:sdfsdfsdf"}, nil).Times(1),
					logger.EXPECT().Println("challenge message received"),
					logger.EXPECT().Println("challenge found"),
				)
			},
			want:    "sdfsdfsdf",
			wantErr: false,
		},
		{
			name: "error_invalid_challenge_header",
			args: args{
				requestID:   "123456",
				requestTime: "1732456988",
			},
			mockCalls: func(conn net.Conn, messenger *mock.MockMessenger, logger *loggerMock.MockLogger) {
				data := []string{"X-Command:Token", "X-Request-id:123456", "X-Request-time:1732456988"}

				gomock.InOrder(
					messenger.EXPECT().Write(conn, data).Return(nil).Times(1),
					logger.EXPECT().Println("token requested"),
					messenger.EXPECT().Read(conn).Return([]string{"X-Chllge:sdfsdfsdf"}, nil).Times(1),
					logger.EXPECT().Println("challenge message received"),
				)
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		messenger := mock.NewMockMessenger(ctrl)
		log := loggerMock.NewMockLogger(ctrl)

		conn, _ := net.Pipe()
		defer conn.Close()

		c := client.NewTCPClient(conn, messenger, log)

		tt.mockCalls(conn, messenger, log)

		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetChallenge(tt.args.requestID, tt.args.requestTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChallenge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetChallenge() got = %v, want %v", got, tt.want)
			}
		})
	}
}
