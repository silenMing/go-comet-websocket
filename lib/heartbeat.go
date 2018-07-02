package lib

import (
	"net"
	"time"
)

func Hearbeat(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case _ = <-readerChannel:
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	case <-time.After(time.Second * 5):
		conn.Close()
	default:
	}
}
