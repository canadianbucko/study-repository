package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/k0kubun/pp"
)

type Message struct {
	person  string
	message string
}

func NewMessage(person string, message string) Message {
	return Message{
		person:  person,
		message: message,
	}
}

func main() {
	fmt.Println("program has been started.")

	iterations := 1 + rand.Intn(9)       // сколько раз в канал записывать будем
	randInterval := 300 + rand.Intn(400) // сколько мы спим

	fmt.Println("всего итераций", iterations)

	myLittleChannel := make(chan Message)

	go func() {
		for i := 0; i <= iterations; i++ {
			convertedIterations := strconv.Itoa(i) // "5"
			convertedInterval := strconv.Itoa(randInterval)
			person := "person number" + convertedIterations
			message := "message" + convertedInterval

			theirMessage := NewMessage(person, message)

			myLittleChannel <- theirMessage

			time.Sleep(time.Duration(randInterval) * (time.Millisecond))

		}
		close(myLittleChannel)
	}()

	for v := range myLittleChannel {
		<-myLittleChannel
		pp.Println(v)
	}
}
