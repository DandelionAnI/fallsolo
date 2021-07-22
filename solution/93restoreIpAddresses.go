package solution

import "strconv"

//dfs回溯 o3^4 o4
func restoreIpAddresses(s string) []string {
	n := len(s)
	parts := make([]int, 4)
	ans := []string{}
	var dfs func(id, index int)
	dfs = func(id, index int) {
		//找到了 4 段 IP 地址并且遍历完了字符串,加入答案
		if id == 4 {
			if index == n {
				ip := ""
				for i := 0; i < 4; i++ {
					ip += strconv.Itoa(parts[i])
					if i != 3 {
						ip += "."
					}
				}
				ans = append(ans, ip)
			}
			return
		}
		// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
		if index == n {
			return
		}
		// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
		if s[index] == '0' {
			parts[id] = 0
			dfs(id+1, index+1)
		}
		// 一般情况，枚举每一种可能性并递归
		add := 0
		for i := index; i < n; i++ {
			add = add*10 + int(s[i]-'0')
			if add > 0 && add <= 255 {
				parts[id] = add
				dfs(id+1, i+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return ans
}
