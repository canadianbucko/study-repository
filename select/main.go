package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanInt := make(chan int)
	chanString := make(chan string)
	ChanFloat64 := make(chan float64)

	go func(ch chan int) {
		for {
			time.Sleep(300 * time.Millisecond)
			ch <- rand.Intn(5)
		}
	}(chanInt)

	go func(ch chan string) {
		for {
			time.Sleep(1 * time.Second)
			ch <- "hello there"
		}
	}(chanString)

	go func(ch chan float64) {
		for {
			time.Sleep(5 * time.Second)
			ch <- 13.7
		}
	}(ChanFloat64)

	for {
		select {
		case int := <-chanInt:
			fmt.Println(int)
		case string := <-chanString:
			fmt.Println(string)
		case float := <-ChanFloat64:
			fmt.Println(float)
		}
	}
}
