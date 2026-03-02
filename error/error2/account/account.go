package account

import (
	"errors"
	"math/rand"
)

type Account interface {
	CashWithdrawal(withdrawal float64) (currentBalance float64, error error)
	CurrentBalance() (currentBalance float64, error error)
	Pay(amount float64, itemDescription string) (currentBalance float64, error error)
}

type bankAccount struct {
	userBalance float64
}

func NewAccount(userBalance float64) Account {
	return &bankAccount{
		userBalance: userBalance,
	}
}
func (b *bankAccount) CashWithdrawal(withdrawal float64) (float64, error) {
	if b.userBalance-withdrawal < 0 {
		err := errors.New("к сожалению, у пользователя недостаточно средств")

		return b.userBalance, err
	}
	b.userBalance = b.userBalance - withdrawal
	return b.userBalance, nil
}

func (b *bankAccount) CurrentBalance() (float64, error) {
	luckiness := rand.Intn(11)
	if luckiness <= 3 {
		err := errors.New("о нет, вам не повезло :( операция не прошла")
		return 0, err
	}
	return b.userBalance, nil
}

func (b *bankAccount) Pay(amount float64, description string) (float64, error) {
	if b.userBalance-amount < 0 {
		err := errors.New("к сожалению, у пользователя недостаточно средств")

		return b.userBalance, err
	}
	luckiness := rand.Intn(11)
	if luckiness <= 3 {
		err := errors.New("о нет, вам не повезло :( операция не прошла")
		return 0, err
	}
	b.userBalance = b.userBalance - amount
	return b.userBalance, nil
}
