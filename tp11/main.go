package main

import (
	"fmt"
	"net"
	"time"
)

func testPort(ip, port string) bool {
	addr := net.JoinHostPort(ip, port)
	timeout := 1 * time.Second

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	addr := "127.0.0.1"

	for i := 1; i <= 65535; i++ {
		if testPort(addr, fmt.Sprintf("%d", i)) {
			fmt.Println(i, "est ouvert")
		}
	}
}
