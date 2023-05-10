package main

import (
	"fmt"
	"os"

	"github.com/kaustuk/go-udp-memcache/pkg/client"
)

func main() {
	conn, err := client.InitUdpConnection()
	if err != nil {
		os.Exit(-1)
	}

	var C client.ClientInterface = conn

	header := []byte{0, 1, 0, 0, 0, 1, 0, 0}
	query := []byte("get key1\r\n")
	fullquery := append(header, query...)

	_, err = C.SendMessage(fullquery)

	if err != nil {
		os.Exit(-1)
	}

	p := make([]byte, 4096)

	nn, err := C.ReceiveMessage(p)

	if err != nil {
		fmt.Printf("Read err %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("%v\n", string(p[:nn]))

	C.CloseConnection()
}
