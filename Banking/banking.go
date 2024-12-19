package Banking

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

// Deposit adds money to the account
func (ba *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive.")
		return
	}
	ba.Balance += amount
	fmt.Printf("%.2f deposited. New balance: %.2f\n", amount, ba.Balance)
}

// Withdraw deducts money if sufficient balance exists
func (ba *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdraw amount must be positive.")
		return
	}
	if ba.Balance < amount {
		fmt.Println("Insufficient balance.")
		return
	}
	ba.Balance -= amount
	fmt.Printf("%.2f withdrawn. New balance: %.2f\n", amount, ba.Balance)
}

// GetBalance prints the current balance
func (ba *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", ba.Balance)
}

// Transaction processes a slice of transactions
// Positive -> Deposit, Negative -> Withdraw
func Transaction(account *BankAccount, transactions []float64) {
	for _, t := range transactions {
		if t > 0 {
			account.Deposit(t)
		} else {
			account.Withdraw(-t)
		}
	}
}
