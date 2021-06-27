package solution

func maxArea(height []int) int {
	l, r, ans := 0, len(height)-1, 0
	for l < r {
		w := r - l
		h := 0
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		cap := w * h
		if cap > ans {
			ans = cap
		}
	}
	return ans
}
