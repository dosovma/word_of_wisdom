package service

type Service struct {
	client TCPClient
	solver Solver
}

func NewService(client TCPClient, solver Solver) *Service {
	return &Service{client: client, solver: solver}
}
