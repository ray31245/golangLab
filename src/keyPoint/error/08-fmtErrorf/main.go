package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := squrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func squrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: squre root of negative number: %v", f)
	}
	return 42, nil
}
