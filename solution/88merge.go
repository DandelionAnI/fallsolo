package solution

func merge1(nums1 []int, m int, nums2 []int, n int)  {
	l, p, r := m-1, m+n-1,n-1
	for r != -1 {
		if l>=0 && nums1[l] > nums2[r] {
			nums1[p] = nums1[l]
			p--
			l--
		} else {
			nums1[p] = nums2[r]
			p--
			r--
		}
	}
}