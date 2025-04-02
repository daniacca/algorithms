package sorting

func merge[T any](left []T, right []T, compare func(a, b T) int) []T {
	out := []T{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if compare(left[i], right[j]) <= 0 {
			out = append(out, left[i])
			i += 1
		} else {
			out = append(out, right[j])
			j += 1
		}
	}

	if i < len(left) {
		out = append(out, left[i:]...)
	}

	if j < len(right) {
		out = append(out, right[j:]...)
	}

	return out
}

func MergeSort[T any](arr []T, compare func(a, b T) int) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid], compare)
	right := MergeSort(arr[mid:], compare)

	return merge(left, right, compare)
}