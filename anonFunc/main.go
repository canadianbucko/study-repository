package main

import (
	"fmt"
)

func main() {
	a := 1
	b := &a
	var result int

	mySecondChannel := make(chan int)
	go func(ch chan int) {
		fmt.Println("anon func 1")
		*b = +1
		ch <- *b

	}(mySecondChannel) // можно и так!
	go func() {
		fmt.Println("anon func 2")
		*b = +1
		mySecondChannel <- *b
	}()
	go func() {
		fmt.Println("anon func 3")
		*b = +1
		mySecondChannel <- *b
	}()

	for i := 0; i <= 2; i++ {
		result += <-mySecondChannel
		fmt.Println(result)
	}
}
