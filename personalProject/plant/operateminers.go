package plant

import (
	"fmt"
	"mill/account"
	"mill/miner"
)

func CreateSmallMiner() int { // возвращает ID созданного шахтера по сути эта функция нунжа чтобы получить экземпляр и положить его в мапу
	defer mtx.Unlock()

	mtx.Lock()

	newID++
	createdMiner := miner.NewSmallMiner() // получаю экземпляр майнера и
	MinerMap[newID] = createdMiner        // сую его в мапу интерфейсов майнеров

	return newID
}

func CreateNormalMiner() int {
	newID++
	createdMiner := miner.NewNormalMiner()
	MinerMap[newID] = createdMiner

	return newID
}

func CreateSuperMiner() int {
	newID++
	createdMiner := miner.NewSuperMiner()
	MinerMap[newID] = createdMiner

	return newID
}

func HireSmallMiner() (int, error) {

	if err := account.SpendMoney(5); err != nil { // списать деньги
		fmt.Println(err)
		return 0, err
	}

	id := CreateSmallMiner() // тут я его создаю

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer mtx.Unlock()
		MinerMap[id].Run(WholePlantContext, id)

		mtx.Lock()
		DeadMiners = append(DeadMiners, id)
		delete(MinerMap, id)
	}()

	return id, nil
}

func HireNormalMiner() (int, error) {
	if err := account.SpendMoney(50); err != nil { // списать деньги
		fmt.Println(err)
		return 0, err
	}

	id := CreateNormalMiner() // тут я его создаю

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer mtx.Unlock()
		MinerMap[id].Run(WholePlantContext, id)

		mtx.Lock()
		DeadMiners = append(DeadMiners, id)
		delete(MinerMap, id)
	}()
	return id, nil
}

func HireSuperMiner() (int, error) {
	if err := account.SpendMoney(450); err != nil { // списать деньги
		fmt.Println(err)
		return 0, err
	}

	id := CreateSuperMiner() // тут я его создаю

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer mtx.Unlock()
		MinerMap[id].Run(WholePlantContext, id)

		mtx.Lock()
		DeadMiners = append(DeadMiners, id)
		delete(MinerMap, id)
	}()
	return id, nil
}
