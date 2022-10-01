package main

import "fmt"

func main() {
	fmt.Println(average(2, 3))
}

func average(xi ...float64) float64 {
	sum := 0.0
	for _, v := range xi {
		sum += v
	}
	average := sum / float64(len(xi))
	return average
}
