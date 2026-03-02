package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
— Мы с вами пишем универсальное хранилище, основанное на мапе
— Ключом в хранилище является строка
— Значением в хранилище является тоже строка
— Необходимо реализовать это хранилище, допуская вероятность конкурентного использования
— Напишите нагрузочные тесты (такие, как мы писали в уроке в теме RWMutex),
 и сравните время выполнения программы с использованием Mutex и с использованием RWMutex
 (p.s. не всегда будет прирост в скорости, не нужно этого пугаться)
*/

var Storage map[string]string = make(map[string]string)
var mtx sync.RWMutex

func Shitting(rand int, convertedRand string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= rand; i++ {
		mtx.Lock()
		Storage[convertedRand] = convertedRand
		mtx.Unlock()
	}

}

func Reading(rand int, convertedRand string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= rand; i++ {
		mtx.RLock()
		_ = Storage[convertedRand]
		mtx.RUnlock()
	}
}

func main() {
	initTime := time.Now()
	wg := &sync.WaitGroup{}
	// rand := rand.Intn(100000)
	rand := 1_000_000
	convertedRand := strconv.Itoa(rand)

	wg.Add(1)
	go Shitting(rand, convertedRand, wg)

	wg.Add(1)
	go Reading(rand, convertedRand, wg)

	wg.Wait()

	fmt.Println(time.Since(initTime))
}
