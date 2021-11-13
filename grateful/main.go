package main

import "grateful/server"

func main() {
	s := server.NewService()
	if err := s.Start("localhost:7777"); err != nil {
		//println(err)
	}
}
