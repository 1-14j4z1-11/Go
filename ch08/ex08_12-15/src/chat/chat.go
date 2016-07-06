package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	channel	chan<- string
	name	string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	autoCloseTime = time.Duration(20 * time.Second)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				go func(c client) {
					c.channel <- msg
				}(cli)
			}

		case cli := <-entering:
			sendCurrenClients(cli, clients)
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 5)
	input := bufio.NewScanner(conn)
	go clientWriter(conn, ch)

	var name string
	ch <- "Your Name >"

	if input.Scan() {
		name = input.Text()
	} else {
		conn.Close()
		return
	}

	closeReset := make(chan struct{})
	go autoClose(conn, closeReset)

	cli := client{channel:ch, name:name}
	ch <- "You are " + cli.name
	messages <- cli.name + " has arrived"
	entering <- cli

	for input.Scan() {
		messages <- cli.name + ": " + input.Text()
		closeReset <- struct{}{}
	}

	leaving <- cli
	messages <- cli.name + " has left"
	conn.Close()
}

func sendCurrenClients(cli client, others map[client]bool) {
	cli.channel <- "<<Current users>>"

	if(len(others) == 0) {
		cli.channel <- "--NONE--"
		return
	}

	for o, _ := range others {
		cli.channel <- "\t" + o.name
	}
}

func autoClose(conn net.Conn, closeReset <-chan struct{}) {
	lastTime := time.Now()
	ticker := time.NewTicker(1 * time.Second)

	for time.Now().Sub(lastTime) < autoCloseTime {
		select {
			case <-closeReset:
				lastTime = time.Now()
			case <-ticker.C:
				// 何もしない
		}
	}

	conn.Close()
	ticker.Stop()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
