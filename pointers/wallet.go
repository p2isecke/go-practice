package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// The Stringer interface is defined in fmt and lets you define how your type is printed when %s is used in format string
type Stringer interface {
	String() string
}

type Bitcoin int
func(b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Use a pointer (*) to the wallet so that we can change the original value
// Change the receiver type from w Wallet --> w *Wallet (pointer to a wallet)
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}