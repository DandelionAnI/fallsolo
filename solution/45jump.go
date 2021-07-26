package solution

//贪心 on o1
func jump(nums []int) int {
	end, maxjump, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxjump = max(maxjump, i+nums[i])
		if i == end {
			end = maxjump
			step++
		}
	}
	return step
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
