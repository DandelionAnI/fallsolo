package solution

import "strconv"

func addStrings(num1 string, num2 string) string {
	ans := ""
	add := 0
	for l, r := len(num1)-1, len(num2)-1; l >= 0 || r >= 0 || add != 0; l, r = l-1, r-1 {
		var x, y int
		if l >= 0 {
			x = int(num1[l] - '0')
		}
		if r >= 0 {
			y = int(num2[r] - '0')
		}
		sum := x + y + add
		add = sum / 10
		ans = strconv.Itoa(sum%10) + ans
	}
	return ans
}
