package main

import "fmt"

type Payment struct{}

func (p Payment) Pay(c CreditCard) {
	c.Pay()
}

type CreditCard struct{}

func (c CreditCard) Pay() {
	fmt.Println("信用卡付款")
}

func main() {
	p := Payment{}
	c := CreditCard{}
	p.Pay(c)
}
