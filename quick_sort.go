// Ascend
func QuickSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	key := arr[begin]
	i, j := begin, end
	for i < j {
		for i < j && arr[j] > key {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] < key {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = key
	QuickSort(arr, begin, i-1)
	QuickSort(arr, i+1, end)
}
