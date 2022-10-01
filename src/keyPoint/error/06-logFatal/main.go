package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happend", err)
		// log.Println("err happend", err)
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("when os.Exit() is call, defrred functions don't run")
}
