package server

import "net"

type TcpServer interface {
	hanle_connection(conn net.Conn)
	Run() error
}
