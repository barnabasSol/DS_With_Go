package binarysearch

func BinarySearch(nums []int, target int) int {
	return binary_search(nums, target, 0, len(nums)-1)

}

func binary_search(nums []int, target, l, r int) int {
	leng := len(nums)
	m := (l + r) / 2
	if r == -1 || l == -1 {
		return -1
	}
	if nums[m] == target {
		return m
	} else {
		if nums[m] < target {
			l = m + 1
			return binary_search(nums, target, l, leng-1)
		} else {
			l = m - 1
			return binary_search(nums, target, l, m-1)
		}
	}
}
