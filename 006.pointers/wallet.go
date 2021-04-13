package pointers

import (
	"errors"
	"fmt"
)

var ErrNotEnoughFund = errors.New("Not enought funds to withdraw")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrNotEnoughFund
	}
	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
