package bubble

func bubbleSort(arr []int) []int {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}
	return arr
}
