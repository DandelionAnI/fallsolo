package solution

//一次遍历 ON o1
func isOneEditDistance(s string, t string) bool {
	n := len(s)
	m := len(t)

	//保证s是较短的字符串
	if n > m {
		return isOneEditDistance(t, s)
	}

	if m-n > 1 {
		return false
	}

	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			if n == m {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}
	return n+1 == m
}
