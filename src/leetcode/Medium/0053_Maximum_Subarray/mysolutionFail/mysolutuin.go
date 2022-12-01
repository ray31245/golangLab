package mysolution

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	sum := nums[0]
	for _, n := range nums[1:] {
		if sum+n >= sum {
			sum = sum + n
		} else {
			sum = 0
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// func sum(nums []int,start,end int) int {
// 	for i := start; i < end; i++ {

// 	}
// }
