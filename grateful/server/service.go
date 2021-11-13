package server

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/kpango/glg"
)

type Service struct {
	srv Server
}

// NewService creates a new Service that handles http connections.
func NewService() *Service {
	return &Service{
		srv: NewServer(),
	}
}

// Start a service that handles grateful shutdown.
func (s *Service) Start(addr string) {
	s.srv.Addr = addr

	// register a goroutine that sends to s.quit on interrupt.
	signal.Notify(s.srv.quit, syscall.SIGINT, syscall.SIGTERM)

	s.srv.Handler = ownHandler()

	_ = glg.Info(fmt.Sprintf("starting server on: %s", addr))

	go func() {
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-s.srv.quit

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// gracefully terminate the server
	if err := s.srv.Shutdown(c); err != nil {
		print(err)
	}
}

// registerHandlers register all handlers for the server.
func ownHandler() http.Handler {
	m := &http.ServeMux{}

	m.Handle("/5sec", http.HandlerFunc(Logging(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		_, _ = w.Write([]byte("done!"))
	})))

	m.Handle("/10sec", http.HandlerFunc(Logging(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		_, _ = w.Write([]byte("done! 10 seconds elapsed"))
	})))

	m.Handle("/15sec", http.HandlerFunc(Logging(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Second)
		_, _ = w.Write([]byte("done! 15 seconds elapsed"))
	})))

	return m
}

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		glg.Info("serving request")
		f(w, r)
		glg.Info("done")
	}
}
