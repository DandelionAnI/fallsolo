package solution

//归并排序 onlogn o1
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	l := arr[0:mid]
	r := arr[mid:]
	return merge2(mergeSort(l), mergeSort(r))
}

func merge2(l []int, r []int) []int {
	var ans []int
	for len(l) != 0 && len(r) != 0 {
		if l[0] <= r[0] {
			ans = append(ans, l[0])
			l = l[1:]
		} else {
			ans = append(ans, r[0])
			r = r[1:]
		}
	}
	for len(l) != 0 {
		ans = append(ans, l[0])
		l = l[1:]
	}
	for len(r) != 0 {
		ans = append(ans, r[0])
		r = r[1:]
	}
	return ans
}
