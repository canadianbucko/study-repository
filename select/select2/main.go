package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hey, my pc is jack shit, so here's the text indicating that the program has been started")
	slice := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	firstGorutinesChannel := make(chan int)
	secondGorutineChannel := make(chan int)
	go func() {
		rand := rand.Intn(len(slice))
		time.Sleep(time.Duration(slice[rand]) * time.Millisecond)
		firstGorutinesChannel <- slice[rand]

	}()

	go func(ch chan int) {
		rand := rand.Intn(len(slice))
		time.Sleep(time.Duration(slice[rand]) * time.Millisecond)
		ch <- slice[rand]
	}(secondGorutineChannel)

	time.Sleep(500 * time.Millisecond)

	select {
	case firstGoroutine := <-firstGorutinesChannel:
		fmt.Println("отработала первая горутина!", firstGoroutine)
	case secondGorotine := <-secondGorutineChannel:
		fmt.Println("отработала вторая горутина", secondGorotine)
	default:
		fmt.Println("не отработала никакая из горутин, вывожу default")
	}
}
