package solution

//dp on o1
func rob2(nums []int) int {
	n := len(nums)
    if n == 1 {
        return nums[0]
    }
    var rob func(a []int) int
	rob = func(a []int) int {
		cur, pre, tmp := 0,0,0
		for i := 0; i < len(a); i++ {
			tmp = cur
			cur = max(cur, pre + a[i])
			pre = tmp
		}
		return cur
	}
	return max(rob(nums[1:]),rob(nums[0:n-1]))
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }