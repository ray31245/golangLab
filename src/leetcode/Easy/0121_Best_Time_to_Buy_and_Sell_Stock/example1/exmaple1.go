package example1

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	max := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
