package net

import (
	"fmt"
	"net"
	"testing"
)

func TestListen(t *testing.T) {
	listener, _ := net.Listen("tcp", "127.0.0.1:3001")
	for {
		conn, _ := listener.Accept()
		go func(conn net.Conn) {
			defer conn.Close()
			var buf [1024]byte
			len, err := conn.Read(buf[:])
			fmt.Println(buf, len, err)
			conn.Write([]byte("i'm server"))
		}(conn)
	}
}
