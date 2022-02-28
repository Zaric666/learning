package sort

func bubbleSort(arr []int) {
	l := len(arr)
	if l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
