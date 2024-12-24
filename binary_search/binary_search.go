package binarysearch

func BinarySearch(nums []int, target int) int {
	return binary_search(nums, target, 0, len(nums)-1)

}

func binary_search(nums []int, target, l, r int) int {
	leng := len(nums)
	m := (l + r) / 2
	if l > r {
		return -1
	}
	if nums[m] == target {
		return m
	} else {
		if nums[m] < target {
			return binary_search(nums, target, m+1, leng-1)
		} else {
			return binary_search(nums, target, m-1, m-1)
		}
	}
}
