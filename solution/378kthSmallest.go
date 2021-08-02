package solution

//二分查找 O(nlog(r−l))，二分查找进行次数为 O(log(r−l))，每次操作时间复杂度为 O(n)
func kthSmallest(matrix [][]int, k int) int {
	n, m, l := len(matrix), len(matrix[0]), matrix[0][0]
	r := matrix[n-1][m-1] + 1
	var cnt func(mid int) int
	cnt = func(mid int) int {
		{
			cnt, j := 0, m-1
			// 每次循环统计比 mid 值小的元素个数
			for i := 0; i < n; i++ {
				// 遍历每行中比 mid 小的元素的个数
				for j >= 0 && mid < matrix[i][j] {
					j--
				}
				cnt += j + 1
			}
			return cnt
		}
	}
	for l < r {
		mid := l + (r-l)/2
		// 如果 count 比 k 小，在大值的那一半继续二分搜索
		if cnt(mid) >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
