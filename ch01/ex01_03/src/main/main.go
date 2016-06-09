package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type action func()

func main() {
	time1 := measurePerformance(printArgs1)
	time2 := measurePerformance(printArgs2)

	fmt.Printf("Normal       : %d [ns]\n", time1.Nanoseconds())
	fmt.Printf("strings.Join : %d [ns]\n", time2.Nanoseconds())
}

func measurePerformance(action action) time.Duration {
	start := time.Now()
	action()

	return time.Since(start)
}

func printArgs1() {

	s, sep := "", ""

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		time.Sleep(1)	// 遅延
	}

	fmt.Println(s)
}

func printArgs2() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
