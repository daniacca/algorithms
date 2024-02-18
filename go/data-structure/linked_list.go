package datastructure

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	head, tail *Node[T]
}

func (l *List[T]) Add(value T) {
	node := &Node[T]{Value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
}

func (l *List[T]) AddArray(values []T) {
	for _, value := range values {
		l.Add(value)
	}
}

func (l *List[T]) Prepend(value T) {
	node := &Node[T]{Value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.Next = l.head
		l.head = node
	}
}

func remove[T any](actual *Node[T], value T, compare func(a, b T) bool) (data T) {
	if actual == nil || actual.Next == nil {
		return
	}

	if compare(actual.Next.Value, value) {
		ret := actual.Next.Value
		actual.Next = actual.Next.Next
		return ret
	}

	return remove(actual.Next, value, compare)
}

func (l *List[T]) RemoveRecursive(value T, compare func(a, b T) bool) (data T) {
	if compare(l.head.Value, value) {
		ret := l.head.Value
		l.head = l.head.Next
		return ret
	}

	return remove(l.head, value, compare)
}

func (l *List[T]) Remove(value T, compare func(a, b T) bool) {
	if l.head == nil {
		return
	}

	if compare(l.head.Value, value) {
		l.head = l.head.Next
		return
	}

	for node := l.head; node.Next != nil; node = node.Next {
		if compare(node.Next.Value, value) {
			if node.Next == l.tail {
				l.tail = node
			}
			node.Next = node.Next.Next
			return
		}
	}
}

func (l *List[T]) ToArray() []T {
	arr := []T{}

	for node := l.head; node != nil; node = node.Next {
		arr = append(arr, node.Value)
	}

	return arr
}

func (l *List[T]) Get(index int) (data T) {
	if l.head == nil {
		return
	}

	node := l.head
	for i := 0; i < index; i++ {
		if node.Next == nil {
			return
		}
		node = node.Next
	}

	return node.Value
}

func (l *List[T]) Set(index int, value T) {
	if l.head == nil {
		return
	}

	node := l.head
	for i := 0; i < index; i++ {
		if node.Next == nil {
			return
		}
		node = node.Next
	}

	node.Value = value
}

func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
}

func (l *List[T]) Find(value T, compare func(a, b T) bool) (data T) {
	if l.head == nil {
		return
	}

	for node := l.head; node != nil; node = node.Next {
		if compare(node.Value, value) {
			return node.Value
		}
	}

	return
}