package sorting

func SelectionSort(arr []int) []int {
	out := []int{}

	for len(arr) > 0 {
		min := arr[0]
		minIndex := 0

		for i := 0; i < len(arr); i++ {
			if arr[i] < min {
				min = arr[i]
				minIndex = i
			}
		}

		out = append(out, min)
		arr = append(arr[:minIndex], arr[minIndex+1:]...)
	}

	return out
}