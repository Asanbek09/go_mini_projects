package server

type Server struct {
	lgr Logger
}

func New(lgr Logger) *Server {
	return &Server{lgr: lgr,}
}

type Logger interface {
	Logf(format string, args ...any)
}