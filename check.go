package main

import (
	"fmt"
	"net"
	"time"
)

func CheckHealth(domain, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, error: %v", address, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable\nFrom: %v\nTo: %v", address, conn.LocalAddr(), conn.RemoteAddr())
	}

	defer conn.Close()

	return status
}
