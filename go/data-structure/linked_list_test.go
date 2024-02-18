package datastructure

import (
	"math/rand"
	"testing"
)

func TestListAdd(t *testing.T) {
	list := List[int]{} // List[int] is a type
	list.Add(1)
	list.Add(2)
	list.Add(3)

	expected := []int{1, 2, 3}
	actual := list.ToArray()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestListAddArray(t *testing.T) {
	list := List[int]{}
	list.AddArray([]int{1, 2, 3})

	expected := []int{1, 2, 3}
	actual := list.ToArray()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestListClear(t *testing.T) {
	list := List[int]{}
	list.AddArray([]int{1, 2, 3})

	list.Clear()

	actual := list.ToArray()

	if len(actual) != 0 {
		t.Errorf("Expected %v, got %v", 0, actual)
	}
}

func TestRemoveRecursive(t *testing.T) {
	list := List[int]{}
	list.AddArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	compare := func(a, b int) bool {
		return a == b
	}

	list.RemoveRecursive(5, compare)

	expected := []int{1, 2, 3, 4, 6, 7, 8, 9}
	actual := list.ToArray()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestRemoveRecursiveBegining(t *testing.T) {
	list := List[int]{}
	list.AddArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	compare := func(a, b int) bool {
		return a == b
	}

	list.RemoveRecursive(1, compare)

	expected := []int{2, 3, 4, 5, 6, 7, 8, 9}
	actual := list.ToArray()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestRemoveRecursiveEnding(t *testing.T) {
	list := List[int]{}
	list.AddArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	compare := func(a, b int) bool {
		return a == b
	}

	list.RemoveRecursive(9, compare)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	actual := list.ToArray()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestRemoveRecursiveIsEmpty(t *testing.T) {
	list := List[int]{} // List[int] is a type

	compare := func(a, b int) bool {
		return a == b
	}

	list.RemoveRecursive(9, compare)
	actual := list.ToArray()

	if len(actual) != 0 {
		t.Errorf("Expected %v, got %v", 0, actual)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

type Complex struct {
	Id string
	Number int
}

func TestRemoveComplex(t *testing.T) {
	list := List[Complex]{}
	data := []Complex{}

	for i := 0; i < rand.Intn(100); i++ {
		data = append(data, Complex{Id: RandStringBytes(10), Number: rand.Intn(100)})
	}

	list.AddArray(data)

	compare := func(a, b Complex) bool {
		return a.Id == b.Id
	}

	toDelete := list.Get(rand.Intn(len(data)))

	deleted := list.RemoveRecursive(Complex{Id: toDelete.Id}, compare)

	if deleted.Id != toDelete.Id {
		t.Errorf("Expected %v, got %v", toDelete, deleted)
	}

	if deleted.Number != toDelete.Number {
		t.Errorf("Expected %v, got %v", toDelete, deleted)
	}

	if len(list.ToArray()) != len(data) - 1 {
		t.Errorf("Expected %v, got %v", 4, len(list.ToArray()))
	}
}