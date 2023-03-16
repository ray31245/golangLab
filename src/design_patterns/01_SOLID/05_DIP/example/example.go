package main

import "fmt"

type Payment struct{}

func (p Payment) Pay(i IPayment) {
	i.Pay()
}

type IPayment interface {
	Pay()
}

type CreditCard struct{}

func (c CreditCard) Pay() {
	fmt.Println("信用卡付款")
}

type QRCode struct{}

func (q QRCode) Pay() {
	fmt.Println("QRCode付款")
}

func main() {
	p := Payment{}
	c := CreditCard{}
	q := QRCode{}
	p.Pay(c)
	p.Pay(q)
}
