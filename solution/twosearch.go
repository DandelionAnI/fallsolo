package solution

func search2(nums []int, target int) int {
	l, r, index := 0, len(nums)-1, -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			index = mid
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return index
}
