package insertion

func insertionSort(arr []int) []int {
	var i, j int

	for i = 1; i < len(arr); i++ {
		tmp := arr[i]
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}
