package server

import (
	"net/http"
	"os"
)

// Server is an abstraction on top of an HTTP server.
type Server struct {
	http.Server

	quit chan os.Signal // used to signal the server a stop signal
}

// NewServer returns an instance of Server.
func NewServer() Server {
	return Server{
		quit: make(chan os.Signal, 1),
	}
}
