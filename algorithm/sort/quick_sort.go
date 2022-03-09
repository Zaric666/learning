package sort

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high) // 分左右两块，和中间位置
		arr = quickSort(arr, low, p-1)     // 对右边再排序
		arr = quickSort(arr, p+1, high)    // 对左边再排序
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] // 取值作为对照
	i := low
	// 小的放左边，大的放右边
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 交换对照值和分块位置，并返回分块位置
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
