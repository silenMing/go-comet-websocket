package main

import (
	"fmt"
	"go-comet-websocket/conf"
	"go-comet-websocket/lib"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {

	//建立socket
	configmap := conf.GetConfig()
	var listen string
	host := conf.GetElemetValue("Host", configmap)
	port := conf.GetElemetValue("Port", configmap)
	timeinterval, err := strconv.Atoi(conf.GetElemetValue("HeartbeatNum", configmap))
	// heartbeatNum := configmap["HeartbeatNum"]
	listen = host + ":" + port
	// log.Print(configmap)
	netListen, err := net.Listen("tcp", listen)
	CheckError(err)
	defer netListen.Close()

	Log("wait client")
	for {
		conn, err := netListen.Accept()

		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " connect success")

		go handleConnect(conn, timeinterval)
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
		Data := (buffer[:n])
		go lib.Hearbeat(conn, messnager, timeout)
		//检测每次是否有数据传过来
		go lib.CheckChannel(Data, messnager)
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
