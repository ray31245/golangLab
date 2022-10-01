package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	recive(eve, odd, quit)
	// fmt.Println("exit")
}

func recive(e, o, q <-chan int) {
	for {
		fmt.Println("-------")
		select {
		case v := <-e:
			fmt.Println("from even channel:", v)
		case v := <-o:
			fmt.Println("from odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
			// fmt.Println("from quit channel:", v)
			// return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	// q <- 0
	close(q)
}
