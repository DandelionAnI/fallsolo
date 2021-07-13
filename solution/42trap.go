package solution

// //dp on on
// func trap(height []int) int {
// 	n := len(height)
// 	if n == 0 {
// 		return 0
// 	}

// 	left := make([]int, n)
// 	left[0] = height[0]
// 	for i := 1; i < n; i++ {
// 		left[i] = max(height[i], left[i-1])
// 	}

// 	right := make([]int, n)
// 	right[n-1] = height[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		right[i] = max(height[i], right[i+1])
// 	}

// 	ans := 0
// 	for i, j := range height {
// 		ans += min(left[i], right[i]) - j
// 	}
// 	return ans
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

//双指针 on o1
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	l, r, ans := 0, n-1, 0
	lmax, rmax := height[l], height[r]
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])
		if lmax < rmax {
			ans += lmax - height[l]
			l++
		} else {
			ans += rmax - height[r]
			r--
		}
	}
	return ans
}