package example3

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

// 單條棧
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
