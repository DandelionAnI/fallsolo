package solution

//堆 onlogn ologn
//建堆的时间代价是O(n)，删除的总代价是O(klogn)，因为k<n，故渐进时间复杂为 O(n+klogn)=O(nlogn)
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, max := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[max] {
		max = l
	}
	if r < heapSize && a[r] > a[max] {
		max = r
	}
	if max != i {
		a[i], a[max] = a[max], a[i]
		maxHeapify(a, max, heapSize)
	}
}
