package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	n     = 30
	procs = 16
)

func main() {
	runtime.GOMAXPROCS(procs)

	start := time.Now()
	ch := make(chan int)
	go fib(n, ch)
	result := <-ch
	end := time.Now()
	fmt.Printf("fib(%d) = %d\tProcs = %d\t Time = %v\n", n, result, procs, end.Sub(start))
}

func fib(n int, result chan<- int) {
	if n <= 1 {
		result <- 1
	} else {
		ch1 := make(chan int)
		ch2 := make(chan int)
		go fib(n-1, ch1)
		go fib(n-2, ch2)
		result <- <-ch1 + <-ch2
	}
}

/*
 * fib(30) = 1346269	Procs = 1	 Time = 6.2163556s
 * fib(30) = 1346269	Procs = 2	 Time = 3.8512203s
 * fib(30) = 1346269	Procs = 4	 Time = 2.4701413s
 * fib(30) = 1346269	Procs = 8	 Time = 1.6880965s
 * fib(30) = 1346269	Procs = 16	 Time = 1.7060975s
 */
