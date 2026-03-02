package main

import (
	"fmt"
	"time"
)

func shit(number int, x *int) {
	/* Я "горутина N, делаю вывод на экран в X раз", где N - номер горутины, а X - номер итерации */
	*x += 1
	fmt.Printf("\n Я горутина %v, делаю вывод на экран в %v раз \n", number, *x)

}

func main() {
	var x int

	/*for i := 1; i <= 5; i++ {
		go shit(i, &x)
	}
	*/
	go shit(1, &x)
	go shit(2, &x)
	go shit(3, &x)
	go shit(4, &x)
	go shit(5, &x)

	time.Sleep(5 * time.Second)
}
