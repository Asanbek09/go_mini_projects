package server

import (
	"fmt"
	"net"
	"context"
	"strconv"

	"google.golang.org/grpc"
	"habits/api"
)

type Server struct {
	api.UnimplementedHabitsServer
	lgr Logger
}

func New(lgr Logger) *Server {
	return &Server{lgr: lgr,}
}

type Logger interface {
	Logf(format string, args ...any)
}

func (s *Server) ListenAndServer(port int) error {
	const addr = "127.0.0.1"

	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	s.lgr.Logf("starting server on port %d\n", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	return nil
}

func (s *Server) CreateHabit(ctx context.Context, req *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	return &api.CreateHabitResponse{}, nil
}