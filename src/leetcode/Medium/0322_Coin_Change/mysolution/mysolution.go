package mysolution

import "fmt"

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	return compute(coins, amount)
}

func compute(coins []int, amount int) int {
	var count int
	var n int
	res := amount + 1
	for _, c := range coins {
		fmt.Println(amount - c)
		n = amount - c
		if n < 0 {
			continue
		} else if n == 0 {
			return 1
		}
		count = compute(coins, amount-c)
		if count > 0 && count < res {
			res = count
		}
	}
	fmt.Println(count)
	if count <= 0 {
		return -1
	}
	fmt.Println("res=", res)
	return res + 1
}
