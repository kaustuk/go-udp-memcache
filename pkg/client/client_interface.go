package client

import (
	"net"
)

type Client struct {
	conn net.Conn
}

type ClientInterface interface {
	CloseConnection() error
	SendMessage([]byte) (int, error)
	ReceiveMessage([]byte) (int, error)
}
