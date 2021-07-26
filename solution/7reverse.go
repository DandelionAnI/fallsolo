package solution

import "math"

func reverse7(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		yu := x % 10
		x /= 10
		ans = ans*10 + yu
	}
	return ans
}
