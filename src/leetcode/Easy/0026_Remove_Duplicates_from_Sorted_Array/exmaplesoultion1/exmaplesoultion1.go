package exmaplesoultion1

import "fmt"

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	lastNum := nums[length-1]
	i := 0
	for i = 0; i < length; i++ {
		if nums[i] == lastNum {
			break
		}
		if nums[i+1] == nums[i] {
			removeElement1(nums, i+1, nums[i])
			fmt.Printf("此時 num = %v length = %v\n", nums, length)
		}
	}
	return i + 1
}

func removeElement1(nums []int, start, val int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
