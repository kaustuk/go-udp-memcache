package main

import (
	"fmt"
	"os"

	"github.com/kaustuk/go-udp-memcache/pkg/client"
)

func main() {
	c, err := client.InitUdpConnection()
	if err != nil {
		os.Exit(-1)
	}

	header := []byte{0, 1, 0, 0, 0, 1, 0, 0}
	query := []byte("get key1\r\n")
	fullquery := append(header, query...)

	_, err = c.SendMessage(fullquery)

	if err != nil {
		os.Exit(-1)
	}

	p := make([]byte, 4096)

	nn, err := c.ReceiveMessage(p)

	if err != nil {
		fmt.Printf("Read err %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("%v\n", string(p[:nn]))

	c.CloseConnection()
}
