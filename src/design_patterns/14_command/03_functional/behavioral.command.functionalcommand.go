package main

import "fmt"

type BankAccount struct {
	Blance int
}

func Deposit(ba *BankAccount, amount int) {
	fmt.Println("Depositing", amount)
	ba.Blance += amount
}

func Withdraw(ba *BankAccount, amount int) {
	if ba.Blance >= amount {
		fmt.Println("Withdrawing", amount)
		ba.Blance -= amount
	}
}

func main() {
	ba := &BankAccount{0}
	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	for _, cmd := range commands {
		cmd()
	}
	fmt.Println(ba)
}
