package methods

import (
	"fmt"
	"math/rand"
)

type Push struct{}

func NewPush() Push {
	return Push{}
}

func (p Push) Send(text string) int {
	fmt.Println("Sending some text using PUSH notification")
	return rand.Int()
}
