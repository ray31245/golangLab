package main

import (
	"fmt"
	"strings"
)

type a struct {
	a string
}

func (a a) printA() {
	fmt.Println(a.a)
}

type b struct {
	a
	b string
}

func (b b) contain() {
	fmt.Println(strings.Contains(b.a.a, b.b))
}

func main() {
	instanceB := b{}
	instanceB.printA()
}
