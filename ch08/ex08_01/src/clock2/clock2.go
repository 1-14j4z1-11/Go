package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if(len(os.Args) <= 1) {
		fmt.Printf("Usage : %s <port>", os.Args[0])
		return
	}

	port, err := strconv.Atoi(os.Args[1])
	if (err != nil) || (port < 0) || (port > 0xFFFF) {
		fmt.Fprintf(os.Stderr, "Invalid port : %s", os.Args[1])
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
