package main

import (
	"Bevis/src/leetcode/Medium/0322_Coin_Change/mysolution"
	"fmt"
)

func main() {
	coins := []int{186, 419, 83, 408}
	amount := 6249
	count := mysolution.CoinChange(coins, amount)
	fmt.Println(count)
}
