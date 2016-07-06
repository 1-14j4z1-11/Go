package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := make(chan string)

	scanner := bufio.NewScanner(c)
	active := true

	finished := func() {
		c.Close()
		close(input)
	}

	go func() {
		for scanner.Scan() {
			input <- scanner.Text()
		}
	}()

	for active {
		ticker := time.NewTicker(10 * time.Second)

		select {
			case <- ticker.C:
				active = false
			case text, ok := <- input:
				if ok {
					go echo(c, text, 1*time.Second)
				} else {
					active = false
				}
		}

		ticker.Stop()
	}

	finished()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
