package wallet

import "errors"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Withdraw(amount float64) error {
	if w.balance-amount < 0 {
		return errors.New("cannot withdraw insufficient wallet balance")
	}
	w.balance -= amount
	return nil
}
