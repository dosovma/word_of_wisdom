package service

// input interface

type IService interface {
	Quote(requestID, requestTime int64) (string, error)
}

// output interfaces

//go:generate mockgen -destination=./mock/tcp_client.go -package=mock . TCPClient
type TCPClient interface {
	GetChallenge(requestID, requestTime string) (string, error)
	GetQuote(token string) (string, error)
	GetTokenBySolution(solution string) (string, error)
}

//go:generate mockgen -destination=./mock/solver.go -package=mock . Solver
type Solver interface {
	Solve(challenge string) (string, error)
}
