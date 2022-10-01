package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//revevie
	bar(c)

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	println(<-c)
}
