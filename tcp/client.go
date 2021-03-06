package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server := "127.0.0.1:8088"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

}

func sender(conn net.Conn) {
	words := "{'id':1}"
	conn.Write([]byte(words))
	fmt.Println("send over")

}
