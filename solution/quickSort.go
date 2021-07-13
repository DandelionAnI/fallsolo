package solution

func QuickSort(arr []int) []int {
	return quickSort1(arr, 0, len(arr)-1)
}

func quickSort1(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort1(arr, left, partitionIndex-1)
		quickSort1(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}
