package plant

import (
	"context"
	"errors"
	"mill/account"
	"mill/miner"
	"sync"
)

// plant владеет всем состояием игры - баланс, покупка, все!
var MinerMap map[int]miner.Miner = make(map[int]miner.Miner)

var DeadMiners []int = make([]int, 0)

var WholePlantContext, WholePlantCancel = context.WithCancel(context.Background())

var wg = &sync.WaitGroup{}

var mtx sync.Mutex

var newID int = 0

func BuyMiner(whichMiner string) (int, error) { // if retuns -1 that means miner WAS NOT CREATED

	switch whichMiner {
	case "small":
		id, err := HireSmallMiner()
		if err != nil {
			return -1, err
		} else {
			return id, nil
		}
	case "normal":
		id, err := HireNormalMiner()
		if err != nil {
			return -1, err
		} else {
			return id, nil
		}
	case "super":
		id, err := HireSuperMiner()
		if err != nil {
			return -1, err
		} else {
			return id, nil
		}
	}
	return -1, errors.New("no argument provided!")
}

func StartPlant() {

	wg.Add(1)
	go account.PassiveIncome(WholePlantContext, wg) // запуск пассивного дохода +1 в сек

	wg.Add(1)
	go account.MinerIncomeOutcome(WholePlantContext, wg) // запуск приема денег=угля

	wg.Wait()

}

func KillPlant() {
	WholePlantCancel()
}

/* шо имеем
StartPlant()
KillPlant()
Println(DeadMiners) --отработавшие майнеры
Println(MinerMap) -- работающие майнеры
Filter("string") структуру + ошибку - работающие + фильтр по классу
BuyPicks() - err
BuyVents() - err
BuyCarts() - err
BoughtEquipment() - возвращает структуру c купленным оборудованием + ценой
account.CurrentBalance() int -  возвращает баланс
MinerInfo(id) - стурктуру с майнером
*/
