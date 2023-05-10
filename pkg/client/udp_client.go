package client

import (
	"fmt"
	"net"
)

func InitUdpConnection() (ClientInterface, error) {
	fmt.Println("*********** init udp connection **********")
	conn, err := net.Dial("udp", "localhost:11212")
	if err != nil {
		return nil, err
	}
	client := Client{conn: conn}
	clientInterface := ClientInterface(&client)
	return clientInterface, nil
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
