package server

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoFunctionality(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer clientConn.Close()

	server := &echoServer{
		host:   "127.0.0.1",
		port:   "8080",
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	go func() {
		server.hanle_connection(serverConn)
	}()

	msg := "Hi, this is your first message\n"
	_, err := clientConn.Write([]byte(msg))
	if err != nil {
		t.Fatal(fmt.Printf("fail writing to client connection due to an error: %v", err))
	}

	readBuff := make([]byte, 512)
	n, err := clientConn.Read(readBuff)
	if err != nil {
		t.Fatal(fmt.Printf("fail reading from client connection due to an error: %v", err))
	}

	assert.Equal(t, string(readBuff[:n]), msg)
}

func TestEchoFunctionality_ReadIgnoreNewline(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer clientConn.Close()

	server := &echoServer{
		host:   "127.0.0.1",
		port:   "8080",
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	go func() {
		server.hanle_connection(serverConn)
	}()

	msg := "Hi, this is your first message\nAnd this is the second batch"
	_, err := clientConn.Write([]byte(msg))
	if err != nil {
		t.Fatal(fmt.Printf("fail writing to client connection due to an error: %v", err))
	}

	readBuff := make([]byte, 512)
	n, err := clientConn.Read(readBuff)
	if err != nil {
		t.Fatal(fmt.Printf("fail reading from client connection due to an error: %v", err))
	}

	assert.Equal(t, msg, string(readBuff[:n]))
}
