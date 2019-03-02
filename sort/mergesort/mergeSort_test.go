package merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	result := mergeSort(arr)
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if reflect.DeepEqual(expect, result) != true {
		t.Errorf("Get %v, expect %v", result, expect)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	for i := 0; i < b.N; i++ {
		mergeSort(arr)
	}
}
