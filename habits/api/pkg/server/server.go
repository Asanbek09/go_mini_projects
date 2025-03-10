package server

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"habits/api"
	"habits/api/pkg/habit"
)

type Server struct {
	api.UnimplementedHabitsServer
	db  Repository
	lgr Logger
}

type Repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	FindAll(ctx context.Context) ([]habit.Habit, error)
}

func New(repo Repository, lgr Logger) *Server {
	return &Server{
		db:  repo,
		lgr: lgr,
	}
}

func (s *Server) ListenAndServe(port int) error {
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

type Logger interface {
	Logf(format string, args ...any)
}