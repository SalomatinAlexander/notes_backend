package server

import (
	"net/http"
	"time"

	"github.com/SalomatinAlexander/noties/internal/handlers"
)

type Server struct {
	h    *handlers.Handler
	port string
}

func New(port string, handler *handlers.Handler) *Server {
	return &Server{
		h:    handler,
		port: port,
	}

}

func (s *Server) Run() error {
	srv := &http.Server{
		Handler:        s.h.InitRout(),
		Addr:           ":" + s.port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return srv.ListenAndServe()

}
