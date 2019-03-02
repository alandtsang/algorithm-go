package merge

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	i, j, k := 0, 0, 0

	tmp := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp[k] = left[i]
			k++
			i++
		} else {
			tmp[k] = right[j]
			k++
			j++
		}
	}

	for i < len(left) {
		tmp[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		tmp[k] = right[j]
		k++
		j++
	}

	return tmp
}
