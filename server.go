package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpSever *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpSever = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpSever.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpSever.Shutdown(ctx)
}
