package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// query := "get key1\r\n"

	// format := []string{"H", "h", "h", "h"}
	// values := []interface{}{1, 0, 1, 0}

	// bp := binary_pack.BinaryPack{}

	// data, err := bp.Pack(format, values)

	// if err != nil {
	// 	panic("error while packing")
	// }

	header := []byte{0, 1, 0, 0, 0, 1, 0, 0}
	query := []byte("get key1\r\n")
	fullquey := append(header, query...)
	fmt.Println("hello from main")
	fmt.Printf("%v\n", fullquey)

	if err := os.WriteFile("file.txt", fullquey, 0644); err != nil {
		log.Fatal(err)
	}

	conn, err := net.Dial("udp", "localhost:11212")

	if err != nil {
		fmt.Printf("Dial err %v", err)
		os.Exit(-1)
	}

	defer conn.Close()

	if _, err = conn.Write(fullquey); err != nil {
		fmt.Printf("Write err %v", err)
		os.Exit(-1)
	}

	p := make([]byte, 4096)
	nn, err := conn.Read(p)
	if err != nil {
		fmt.Printf("Read err %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("%v\n", string(p[:nn]))

	// fmt.Printf("%v\n", query)
}
