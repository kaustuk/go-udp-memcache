package client

import (
	"fmt"
	"net"
)

func InitUdpConnection() (*Client, error) {
	fmt.Println("*********** init udp connection **********")
	conn, err := net.Dial("udp", "localhost:11212")
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (client *Client) CloseConnection() error {
	fmt.Println("************* close connection *************")
	return client.conn.Close()
}

func (client *Client) SendMessage(b []byte) (int, error) {
	return client.conn.Write(b)
}

func (client *Client) ReceiveMessage(b []byte) (int, error) {
	return client.conn.Read(b)
}
