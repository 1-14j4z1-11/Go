package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if(len(os.Args) <= 1) {
		fmt.Printf("Usage : %s <server...>", os.Args[0])
		return
	}

	done := make(chan struct{})

	for _, arg := range os.Args[1:] {
		go startConn(arg, done)
	}

	for range os.Args[1:] {
		<- done
	}
}

func startConn(server string, done chan <- struct{}) {
	fmt.Printf("Start connection : %s\n", server)
	defer func() {
		done <- struct{}{}
	}()

	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
