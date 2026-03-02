package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var vote atomic.Int64

func AddVote(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		vote.Add(1)
	}
}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(10)
	go AddVote(wg)
	go AddVote(wg)
	go AddVote(wg)
	go AddVote(wg)
	go AddVote(wg)

	go AddVote(wg)
	go AddVote(wg)
	go AddVote(wg)
	go AddVote(wg)
	go AddVote(wg)
	wg.Wait()

	fmt.Println(vote.Load())

}
