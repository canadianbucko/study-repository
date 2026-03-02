package miner

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

/*
- Сильный шахтёр:
  - Оплата труда: 450 угля
  - Энергия: хватит на 60 добыч угля
  - Получает угля за одну добычу: 10
  - Перерыв между добычами: 1 секнда
  - С каждой новой добычей, характеристика "Получает угля за одну добычу" увеличивается на 3 единицы
*/
type SuperMiner struct {
	cost   int
	energy int
	income int
	rest   int
}

func NewSuperMiner() *SuperMiner {
	return &SuperMiner{
		cost:   450,
		energy: 60,
		income: 10,
		rest:   1,
	}
}

func (s *SuperMiner) Run(ctx context.Context, id int) {
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

			s.income += 3

		}

	}
	msg := "Майнер " + strconv.Itoa(id) + " завершает свою работу, истощен! "
	fmt.Println(msg)
}

func (s *SuperMiner) Info() *MinerInfo {
	return &MinerInfo{
		Class:  "super",
		Cost:   s.cost,
		Energy: s.energy,
		Income: s.income,
		Rest:   s.rest,
	}
}
