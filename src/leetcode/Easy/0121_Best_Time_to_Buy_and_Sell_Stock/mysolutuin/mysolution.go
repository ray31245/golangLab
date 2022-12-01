package mysolutuin

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min := prices[0]
	temp := 0
	max := 0
	for _, p := range prices[1:] {
		if p-min > temp {
			temp = p - min
		}
		if temp > max {
			max = temp
		}
		if p < min {
			min = p
		}
	}
	return max
}
