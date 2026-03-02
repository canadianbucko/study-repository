package account

import (
	"context"
	"errors"
	"fmt"
	"mill/miner"
	"strconv"
	"sync"
	"time"
)

var Money int = 100000 // базовая сумма денег для начала игры

var mtx sync.Mutex

func PassiveIncome(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping plant's passive income!")
			return
		default:
			time.Sleep(1 * time.Second)
			mtx.Lock()
			Money += 1
			mtx.Unlock()
		}
	}

}

var Multiplier int = 1

func MinerIncomeOutcome(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case dengi := <-miner.CoalTransferPoint:
			mtx.Lock()
			Money += dengi * Multiplier
			mtx.Unlock()
		case <-ctx.Done():
			fmt.Println("stopping miners income/outcome!")
			return
		}
	}

}

func SpendMoney(amount int) error {
	defer mtx.Unlock()

	mtx.Lock()
	if Money-amount >= 0 {
		Money -= amount
		return nil
	}

	return errors.New("insufficient funds")

}

func CurrentBalance() string { //текуший баланс

	mtx.Lock()
	balance := Money
	mtx.Unlock()

	msg := "your current balance is " + strconv.Itoa(balance)

	return msg

}
