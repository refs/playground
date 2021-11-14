package main

import "grateful/server"

func main() {
	s := server.NewService()
	s.Start("localhost:7777")
}
