package solution

import "sort"

//排序+双指针 onlogn ologn
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})

	ans := [][]int{}
	pre := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		//如果右边界比下一组左边界小，则将此组加入ans
		if pre[1] < cur[0] {
			ans = append(ans, pre)
			pre = cur
		} else {
			//如果右边界比下一组左边界大，则合并此组
			pre[1] = max(pre[1], cur[1])
		}
	}
	ans = append(ans, pre)
	return ans
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }