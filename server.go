package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	//建立socket
	netListen, err := net.Listen("tcp", "127.0.0.1:9001")
	CheckError(err)
	defer netListen.Close()

	Log("wait client")

	for {
		conn, err := netListen.Accept()

		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " connect success")

		go handleConnect(conn)
	}
}

// 处理链接

func handleConnect(conn net.Conn, timeout int) {
	buffer := make([]byte, 2048)
	messnager := make(chan byte)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}
