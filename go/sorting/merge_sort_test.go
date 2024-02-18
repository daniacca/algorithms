package sorting

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{1, 5, 3, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	actual := MergeSort(arr, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestMergeSortEmpty(t *testing.T) {
	arr := []int{}
	expected := []int{}
	actual := MergeSort(arr, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestMergeSortOne(t *testing.T) {
	arr := []int{1}
	expected := []int{1}
	actual := MergeSort(arr, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestMergeSortTwo(t *testing.T) {
	arr := []int{2, 1}
	expected := []int{1, 2}
	actual := MergeSort(arr, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestMergeSortBigRandomArray(t *testing.T) { // 100k elements
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	actual := MergeSort(arr, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(arr)-1; i++ {
		if actual[i] > actual[i+1] {
			t.Errorf("Index %d and %d are not ordered", i, i+1)
		}
	}
}

func TestMergeSortBigReversedArray(t *testing.T) { // 100k elements
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}

	expected := []int{}
	for i := 1; i <= 100000; i++ {
		expected = append(expected, i)
	}

	actual := MergeSort(arr, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	for i := 0; i < b.N; i++ {
		MergeSort(arr, func(a, b int) int {
			return a - b
		})
	}
}