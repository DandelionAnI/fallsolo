package solution

//滑动窗口 hash C字符集大小 o(C*s+t) o(C)
func minWindow(s string, t string) string {
	tsub, ssub := map[byte]int{}, map[byte]int{}
	for i := range t {
		tsub[t[i]]++
	}
	l, r, cnt, minw := 0, -1, 0, len(s)+1
	lf, rf := -1, -1
	for l < len(s) {
		if r+1 < len(s) && cnt < len(t) {
			ssub[s[r+1]]++
			if ssub[s[r+1]] <= tsub[s[r+1]] {
				cnt++
			}
			r++
		} else {
			if r-l+1 < minw && cnt == len(t) {
				minw = r - l + 1
				lf = l
				rf = r
			}
			if ssub[s[l]] == tsub[s[l]] {
				cnt--
			}
			ssub[s[l]]--
			l++
		}
	}
	ans := ""
	if lf != -1 {
		ans = s[lf:rf+1]
	}
	return ans
}
