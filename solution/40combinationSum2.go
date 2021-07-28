package solution

import "sort"

//先排序为了去重，dfs回溯 o(n*n!) on
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	a := []int{}
	ans := [][]int{}

	var dfs func(target, index int)
	dfs = func(target, index int) {
		if target == 0 {
			tmp := make([]int, len(a))
			copy(tmp, a)
			ans = append(ans, tmp)
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			if target >= candidates[i] {
				a = append(a, candidates[i])
				dfs(target-candidates[i], i+1)
				a = a[:len(a)-1]
			}
		}
	}
	dfs(target, 0)
	return ans
}
