package sorting

import (
	"math/rand"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 5, 3, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	actual := SelectionSort(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestSelectionSortEmpty(t *testing.T) {
	arr := []int{}
	expected := []int{}
	actual := SelectionSort(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestSelectionSortOne(t *testing.T) {
	arr := []int{1}
	expected := []int{1}
	actual := SelectionSort(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func TestSelectionSortTwo(t *testing.T) {
	arr := []int{2, 1}
	expected := []int{1, 2}
	actual := SelectionSort(arr)

	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual[i])
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100000))
	}
	
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}