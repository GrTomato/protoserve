package main

import (
	"fmt"
	"grtomato/protoserve/server"
	"os"
)

func main() {
	logger := NewLogger()

	host := os.Args[1]
	port := os.Args[2]

	logger.Debug(fmt.Sprintf("Input host = %v, port = %v", host, port))

	echoServer := server.NewEchoServer(host, port, logger)
	err := echoServer.Run()
	if err != nil {
		logger.Warn(fmt.Sprintf("The server was stopped with error: %v", err))
		os.Exit(1)
	}
}
