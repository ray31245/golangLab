package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recive
	cs := make(chan<- int) //send
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
}
