package main

import "grateful/server"

func main() {
	s := server.NewService()
	s.Start("localhost:7777")
}

//package main
//
//import (
//	"context"
//	"log"
//	"net/http"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//func main() {
//	hs, logger := setup()
//
//	go func() {
//		logger.Printf("Listening on http://0.0.0.0%s\n", hs.Addr)
//
//		if err := hs.ListenAndServe(); err != http.ErrServerClosed {
//			logger.Fatal(err)
//		}
//	}()
//
//	graceful(hs, logger, 5*time.Second)
//}
//
//func setup() (*http.Server, *log.Logger) {
//	addr := ":" + os.Getenv("PORT")
//	if addr == ":" {
//		addr = ":2017"
//	}
//
//	hs := &http.Server{Addr: addr, Handler: &server{}}
//
//	return hs, log.New(os.Stdout, "", 0)
//}
//
//func graceful(hs *http.Server, logger *log.Logger, timeout time.Duration) {
//	stop := make(chan os.Signal, 1)
//
//	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
//
//	<-stop
//
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//
//	logger.Printf("\nShutdown with timeout: %s\n", timeout)
//
//	if err := hs.Shutdown(ctx); err != nil {
//		logger.Printf("Error: %v\n", err)
//	} else {
//		logger.Println("Server stopped")
//	}
//}
//
//type server struct{}
//
//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	time.Sleep(5 * time.Second)
//	w.Write([]byte("Hello, World!"))
//}
