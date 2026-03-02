package plant

import (
	"errors"
	"mill/account"
)

var picksBought bool = false
var ventsBought bool = false
var cartsBought bool = false

type equipDTO struct {
	Bought bool `json:"bought"`
	Price  int  `json:"price"`
}

func NewEquipDTO(which string) equipDTO {
	defer mtx.Unlock()
	mtx.Lock()

	switch which {
	case "picks":
		return equipDTO{
			Bought: picksBought,
			Price:  3000,
		}
	case "vents":
		return equipDTO{
			Bought: ventsBought,
			Price:  15000,
		}
	case "carts":
		return equipDTO{
			Bought: cartsBought,
			Price:  50000,
		}

	}
	return equipDTO{}
}

func BuyPicks() error {
	defer mtx.Unlock()

	mtx.Lock()

	if picksBought == false {
		err := account.SpendMoney(3000)
		if err != nil {
			return err
		}
		account.Multiplier += 1
		picksBought = true
	}
	return errors.New("oops! picks have been already purschased!")
}

func BuyVents() error {
	defer mtx.Unlock()

	mtx.Lock()

	if ventsBought == false {
		err := account.SpendMoney(15000)
		if err != nil {
			return err
		}
		account.Multiplier += 2
		ventsBought = true
	}
	return errors.New("oops! vents have been already purschased!")
}

func BuyCarts() error {
	defer mtx.Unlock()

	mtx.Lock()

	if cartsBought == false {
		err := account.SpendMoney(50000)
		if err != nil {
			return err
		}
		account.Multiplier += 4
		cartsBought = true
	}
	return errors.New("oops! carts have been already purschased!")
}
