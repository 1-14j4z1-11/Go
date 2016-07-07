package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})

	go rally(ch1, ch2, done)
	go rally(ch2, ch1, done)

	fmt.Printf("[%v]\tStart\n", time.Now())
	ch1 <- 0
	go func() {
		time.Sleep(time.Duration(1 * time.Second))
		close(done)
	}()

	<-done
	var count int

	select {
	case count = <-ch1:
	case count = <-ch2:
	}

	fmt.Printf("[%v]\tEnd  Count = %d\n", time.Now(), count)
}

func rally(send chan<- int, receive <-chan int, done <-chan struct{}) {
	count := -1
	for {
		select {
		case <-done:
			return
		case count = <-receive:
			send <- count + 1
		}
	}
}
