package solution

//贪心 on o1
func canJump(nums []int) bool {
	maxjump := 0
	for i, num := range nums {
		if i > maxjump {
			return false
		}
		maxjump = max(maxjump, i+num)
	}
	return true
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
