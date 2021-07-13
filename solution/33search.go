package solution

//二分 ologn o1
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (r-l)/2 +l
		if target == nums[mid] {
			return mid
		}
		if nums[mid] <= nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	} 
	return -1
}