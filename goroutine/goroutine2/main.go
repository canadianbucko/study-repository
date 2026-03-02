package main

import (
	"fmt"
	"time"
)

func shit(myFirstChannel chan int) {
	rand := 1
	myFirstChannel <- rand
	fmt.Println("передал значение, засыпаю!")
	time.Sleep(50000 * time.Second)
}

func main() {
	myFirstChannel := make(chan int, 5)
	fmt.Println("start")
	var result int
	for i := 1; i <= 5; i++ {
		go shit(myFirstChannel)
		result += <-myFirstChannel
		fmt.Println(result)

	}

}
