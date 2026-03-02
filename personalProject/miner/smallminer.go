package miner

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type SmallMiner struct {
	cost   int
	energy int
	income int
	rest   int
}

func NewSmallMiner() *SmallMiner { // просто делает конкретный экземпляр майнера (конструктор)
	return &SmallMiner{
		cost:   5,
		energy: 30,
		income: 1,
		rest:   3, // seconds
	}

}

func (s *SmallMiner) Run(ctx context.Context, id int) {
	for s.energy > 0 { // пока не кончится энергия
		select {
		case <-ctx.Done():
			msg := "Майнер " + strconv.Itoa(id) + " завершает свою работу!"
			fmt.Println(msg) // и потом еще запихать это куда-нибудь в ответ http
			return
		default:
			// пока не кончится энергия
			s.energy -= 1
			CoalTransferPoint <- s.income                        // отдаем значение в канал
			time.Sleep(time.Duration(s.rest) * time.Microsecond) // rest 3

		}

	}
	msg := "Майнер " + strconv.Itoa(id) + " завершает свою работу, истощен! "
	fmt.Println(msg)
}

func (s *SmallMiner) Info() *MinerInfo {
	return &MinerInfo{
		Class:  "small",
		Cost:   s.cost,
		Energy: s.energy,
		Income: s.income,
		Rest:   s.rest,
	}
}
