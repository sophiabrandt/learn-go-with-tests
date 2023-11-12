package fintech

import (
	"fmt"
)

// walletError is a custom error type.
type walletError string

// Error is a built-in interface that returns a string.
func (we walletError) Error() string {
	return string(we)
}

// ErrInssuficientFunds shows an error message when the wallet has insufficient funds.
const ErrInssuficientFunds = walletError("cannot withdraw, insufficient funds")

// Bitcoin represents the bitcoin inside a wallet.
type Bitcoin int

// String() is a built-in interface that returns a string.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet manages a Bitcoin wallet.
type Wallet struct {
	balance Bitcoin
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	// same as (*w).balance += amount
	w.balance += amount
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw subtracts some Bitcoin from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInssuficientFunds
	}
	w.balance -= amount
	return nil
}
