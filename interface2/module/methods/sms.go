package methods

import (
	"fmt"
	"math/rand"
)

type Sms struct{}

func NewSms() Sms {
	return Sms{}
}

func (p Push) Sms(text string) int {
	fmt.Println("Sending some text using SMS-message")
	return rand.Int()
}
