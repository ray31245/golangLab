package example1

import "fmt"

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		fmt.Println("amount=", i)
		for j := 0; j < len(coins); j++ {
			fmt.Println("coin=", coins[j])
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
				fmt.Println(dp)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
