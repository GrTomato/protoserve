package server

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
)

type echoServer struct {
	host   string
	port   string
	logger *slog.Logger
}

func NewEchoServer(host string, port string, logger *slog.Logger) TcpServer {
	return &echoServer{
		host:   host,
		port:   port,
		logger: logger,
	}
}

func (s *echoServer) Run() error {
	s.logger.Debug("A new echo server started")

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", s.host, s.port))
	if err != nil {
		s.logger.Error(fmt.Sprintf("unable to create listener: %v", err))
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		s.logger.Debug(fmt.Sprintf("A new connection was accepted from: %v", conn.RemoteAddr()))

		if err != nil {
			s.logger.Error(fmt.Sprintf("unable to obtain next connection: %v", err))
			return err
		}

		go s.hanle_connection(conn)
	}
}

func (s *echoServer) hanle_connection(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 4096) // 4kb
	for {
		// read from connection
		nRead, err := conn.Read(buff)

		if nRead > 0 {
			// write back exactly the number of writtent bytes
			nWritten, err := conn.Write(buff[0:nRead])
			if err != nil {
				s.logger.Error(fmt.Sprintf("remote: %v, unable to write to target %v\n", conn.RemoteAddr(), err))
				break
			}

			if nRead != nWritten {
				s.logger.Warn(fmt.Sprintf("Warning: the number of read bytes do not equal to the number of written bytes: read=%v, written=%v", nRead, nWritten))
			}

			s.logger.Debug(fmt.Sprintf("Bytes written: %v", nWritten))
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				s.logger.Debug(fmt.Sprintf("Remote: %v. Closing connection with remote due to an EOF\n", conn.RemoteAddr()))
				break
			}
			s.logger.Debug(fmt.Sprintf("remote: %v, unable to read a line, closing connection: %v\n", conn.RemoteAddr(), err))
			break
		}
	}
}
