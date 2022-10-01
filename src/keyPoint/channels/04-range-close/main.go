package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//recive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("exit")
}

func foo(c chan<- int) {
	// c <- 42
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
