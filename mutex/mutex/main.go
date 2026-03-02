package main

import (
	"fmt"
	"sync"
)

var slice []int
var mtx sync.Mutex

func AddVote(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
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

	fmt.Println(len(slice))

}
