package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := squrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func squrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math: squre root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W ", nme}
	}
	return 42, nil
}
