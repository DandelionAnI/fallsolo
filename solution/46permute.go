package solution

//回溯 o(n*n!) on
func permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	used := make([]bool, n)
	var add func(meta []int)
	add = func(meta []int) {
		if len(meta) == n {
			// 为什么加入解集时，要将数组（描述中数组都是切片）内容复制到一个新的数组里，再加入解集?
			// 这个 track 变量是一个地址引用，结束当前递归，将它加入 res，后续的递归分支还要继续进行搜
			// 索，还要继续传递这个 track ，这个地址引用所指向的内存空间还要继续被操作，所以 res 中的
			// track 所引用的内容会被改变，这就造成了 res 中的内容随 track 变化。
			// 所以要复制 track 内容到一个新的数组里，然后放入 res，这样后续对 track 的操作，
			// 就不会影响已经放入 res 的内容。
			tmp := make([]int, n)
			copy(tmp, meta)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				meta = append(meta, nums[i])
				add(meta)
				meta = meta[:len(meta)-1]
				used[i] = false
			}
		}
	}
	meta := []int{}
	add(meta)
	return ans
}
