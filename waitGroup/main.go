package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker1(wg *sync.WaitGroup) {
	defer wg.Done()
	rand := 500 + rand.Intn(500)
	time.Sleep(time.Duration(rand) * (time.Millisecond))
	fmt.Println("I am a worker and I'm watering some stuff in here!")

}

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	rand := 500 + rand.Intn(500)
	time.Sleep(time.Duration(rand) * (time.Millisecond))
	fmt.Println("I am a worker and I'm watering some stuff in here!")

}

func worker3(wg *sync.WaitGroup) {
	defer wg.Done()
	rand := 500 + rand.Intn(500)
	time.Sleep(time.Duration(rand) * (time.Millisecond))
	fmt.Println("I am a worker and I'm watering some stuff in here!")

}
func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker1(wg)

	wg.Add(1)
	go worker2(wg)

	wg.Add(1)
	go worker2(wg)

	wg.Wait()

}
