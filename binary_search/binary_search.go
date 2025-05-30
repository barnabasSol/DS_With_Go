package binarysearch

func BinarySearch(nums []int, target int) int {
	return binary_search(nums, target, 0, len(nums)-1)

}

func binary_search(nums []int, target, l, r int) int {
	if l > r {
		return -1
	}
	m := (r + l) / 2
	if nums[m] == target {
		return m
	} else {
		if nums[m] < target {
			return binary_search(nums, target, m+1, r)
		} else {
			return binary_search(nums, target, l, m-1)
		}
	}
}
