package solution

//双指针 on on
func reverseWords(s string) string {
	ans := ""
	b := []byte(s)
	//字符串逐字反转 ab cd ef -> fe dc ba
	var reverse func(l, r int)
	reverse = func(l, r int) {
		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}
	reverse(0, len(b)-1)
	//遍历，再把每一个单词翻转 fe dc ba -> ef cd ab
	l, r, i := 0, 0, 0
	for i < len(b) {
		if b[i] != ' ' {
			l = i
			for i < len(b) && b[i] != ' ' {
				i += 1
			}
			r = i - 1
			reverse(l, r)
			ans = ans + " " + string(b[l:r+1])
		}
		i += 1
	}
	return ans[1:]
}
