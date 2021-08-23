package solution

func numRabbits(ans []int) int {
	sum, m := 0, make(map[int]int)
	for _, v := range ans {
		if m[v] == 0 {
			m[v] += v
			sum += v + 1
		} else {
			m[v]--
		}
	}
	return sum
}
