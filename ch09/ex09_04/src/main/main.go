package main

import (
	"fmt"
	"time"
)

const (
	total = 1000000
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	fmt.Printf("[%v] Start\n", time.Now())
	go pipeLine(total, ch, done)

	ch <- 0
	<-done
}

func pipeLine(count int, ch <-chan int, done chan<- struct{}) {
	if count == 0 {
		x := <-ch
		fmt.Printf("[%v] End  Created = %d\n", time.Now(), x)
		done <- struct{}{}
	} else {
		ch1 := make(chan int)
		go pipeLine(count-1, ch1, done)
		x := <-ch
		ch1 <- (x + 1)
	}
}
