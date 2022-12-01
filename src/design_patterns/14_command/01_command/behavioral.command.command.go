package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	blance int
}

func (b *BankAccount) Deposit(amount int) {
	b.blance += amount
	fmt.Println("Deposited", amount, "\b, blance is now", b.blance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.blance-amount >= overdraftLimit {
		b.blance -= amount
		fmt.Println("Withdrew", amount, "\b, blance is now", b.blance)
		return true
	}
	return false
}

// type Command interface {
// 	Call()
// 	Undo()
// }

type Action int

const (
	Deposite Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account  *BankAccount
	action   Action
	amount   int
	succeded bool
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposite:
		b.account.Deposit(b.amount)
		b.succeded = true
	case Withdraw:
		b.succeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeded {
		return
	}
	switch b.action {
	case Deposite:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposite, 100)
	cmd.Call()
	cmd.Call()
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	cmd2.Call()
	fmt.Println(ba)
}
