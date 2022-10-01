package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
	c2 := make(chan int, 2)
	c2 <- 43
	c2 <- 45
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
