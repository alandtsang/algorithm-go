package shell

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	result := shellSort(arr)
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if reflect.DeepEqual(expect, result) != true {
		t.Errorf("Get %v, expect %v", result, expect)
	}
}

func BenchmarkShellSort(b *testing.B) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	for i := 0; i < b.N; i++ {
		shellSort(arr)
	}
}
