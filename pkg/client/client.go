package client

import "net"

type Client struct {
	conn net.Conn
}
