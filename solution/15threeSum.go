package solution

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	//枚举第一个
	for a := 0; a < n; a++ {
		//a需要和上一次枚举的数不相同
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		c := n - 1
		target := -nums[a]
		// 枚举 b
		for b := a + 1; b < n; b++ {
			//b需要和上一次枚举的数不相同
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			//需要保证 b 的指针在 c 的指针的左侧
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}
