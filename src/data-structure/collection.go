package datastructure

type Collection[T any] interface {
	Add(value T)
	AddArray(values []T)
	Clear()
	Remove(value T, compare func(a, b T) bool)
	ToArray() []T
	Get(index int) T
	Set(index int, value T)
	Find(value T, compare func(a, b T) bool) T
}