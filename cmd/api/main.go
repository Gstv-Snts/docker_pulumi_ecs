package main

import (
	"ecs_pulumi_go/internal/server"
	"fmt"
)

func main() {
	println("opening the server")

	server := server.NewServer()

	println("server created")

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
