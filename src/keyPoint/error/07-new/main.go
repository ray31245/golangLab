package main

import (
	"errors"
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
		return 0, errors.New("norgate math: squre root of negative number")
	}
	return 42, nil
}
