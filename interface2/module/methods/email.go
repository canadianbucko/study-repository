package methods

import (
	"fmt"
	"math/rand"
)

type Email struct{}

func NewEmail() Email {
	return Email{}
}

func (p Email) Send(text []string) int {
	fmt.Println("Sending some text using EMAIL")
	return rand.Int()
}
