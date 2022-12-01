package mysolution

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := m - 1
	k := n - 1
	for i := len(nums1) - 1; k >= 0 && j >= 0 && i >= 0; i-- {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
	}
	for ; k >= 0; k-- {
		nums1[k] = nums2[k]
	}
}
