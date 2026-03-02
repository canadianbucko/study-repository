package miner

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type NormalMiner struct {
	cost   int
	energy int
	income int
	rest   int
}

func NewNormalMiner() *NormalMiner {
	return &NormalMiner{
		cost:   50,
		energy: 45,
		income: 3,
		rest:   2,
	}
}

func (s *NormalMiner) Run(ctx context.Context, id int) {
	for s.energy > 0 { // пока не кончится энергия
		select {
		case <-ctx.Done():
			msg := "Майнер " + strconv.Itoa(id) + " завершает свою работу!"
			fmt.Println(msg) // и потом еще запихать это куда-нибудь в ответ http
			return
		default:
			// пока не кончится энергия

			s.energy -= 1

			CoalTransferPoint <- s.income                   // отдаем значение в канал
			time.Sleep(time.Duration(s.rest) * time.Second) // rest 2

		}

	}
	msg := "Майнер " + strconv.Itoa(id) + " завершает свою работу, истощен! "
	fmt.Println(msg)
}

func (s *NormalMiner) Info() *MinerInfo {
	return &MinerInfo{
		Class:  "normal",
		Cost:   s.cost,
		Energy: s.energy,
		Income: s.income,
		Rest:   s.rest,
	}
}
