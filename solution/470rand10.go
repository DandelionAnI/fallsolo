package solution

import "math/rand"

func rand10() int {
	rand40 := 40
	for rand40 >= 40 {
		rand40 = rand.Intn(7)*7 + rand.Intn(7)
	}
	return rand40%10 + 1
}

// rand.Intn(7) 等概率地产生 [0,6].
// rand.Intn(7) *7 等概率地产生 0, 7, 14, 21, 28, 35, 42
// rand.Intn(7) * 7 + rand.Intn(7) 等概率地产生 [0, 48] 这 49 个数字
// 如果步骤 4 的结果大于等于 40，那么就重复步骤 4，直到产生的数小于 40
// 把步骤 5 的结果 mod 10 再加 1， 就会等概率的随机生成 [1, 10]
