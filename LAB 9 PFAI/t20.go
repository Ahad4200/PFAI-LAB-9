package main

import (
	"errors"
	"fmt"
)

// BankAccount struct with balance
type BankAccount struct {
	balance float64
}

// Deposit adds amount to balance
func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// Withdraw subtracts amount from balance
func (a *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

// Balance returns the current balance
func (a *BankAccount) Balance() float64 {
	return a.balance
}

func main() {
	// Create a new account
	account := BankAccount{balance: 100}
	fmt.Printf("Initial balance: $%.2f\n", account.Balance())

	// Test deposit
	err := account.Deposit(50)
	if err != nil {
		fmt.Printf("Deposit error: %s\n", err)
	} else {
		fmt.Printf("After deposit: $%.2f\n", account.Balance())
	}

	// Test successful withdrawal
	err = account.Withdraw(30)
	if err != nil {
		fmt.Printf("Withdrawal error: %s\n", err)
	} else {
		fmt.Printf("After withdrawal: $%.2f\n", account.Balance())
	}

	// Test overdraft prevention
	err = account.Withdraw(200)
	if err != nil {
		fmt.Printf("Withdrawal error: %s\n", err)
	} else {
		fmt.Printf("After withdrawal: $%.2f\n", account.Balance())
	}
}
