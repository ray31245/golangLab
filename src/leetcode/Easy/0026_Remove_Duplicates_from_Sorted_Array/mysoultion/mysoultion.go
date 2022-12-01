package mysoultion

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	curr := 0
	for i, n := range nums[1:] {
		if n > nums[i] {
			curr++
			nums[curr] = n
		}
	}
	return curr + 1
}
