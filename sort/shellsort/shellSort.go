package shell

func shellSort(arr []int) []int {
	var gap int
	var i, j int

	n := len(arr)
	for gap = n / 2; gap > 0; gap /= 2 {
		for i = gap; i < n; i++ {
			tmp := arr[i]
			for j = i; j >= gap && arr[j-gap] > tmp; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = tmp
		}
	}
	return arr
}
