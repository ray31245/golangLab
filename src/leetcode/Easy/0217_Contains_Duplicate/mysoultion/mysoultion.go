package mysoultion

func ContainsDuplicate(nums []int) bool {
	return containsDuplicate(nums)
}

func containsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))
	for _, n := range nums {
		if exist := record[n]; exist {
			return true
		}
		record[n] = true
	}
	return false
}
