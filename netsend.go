package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage:")
		fmt.Println("  netsend <ip> <port> \"message\"")
		os.Exit(1)
	}

	ip := os.Args[1]
	port := os.Args[2]
	message := os.Args[3]

	if _, err := strconv.Atoi(port); err != nil {
		fmt.Println("Invalid port.")
		os.Exit(1)
	}

	address := net.JoinHostPort(ip, port)

	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Send failed:", err)
		os.Exit(1)
	}

	fmt.Println("Sent successfully.")
}