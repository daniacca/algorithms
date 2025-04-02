package integer_multiplication

import "math"

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func SchoolMultiply(x int, y int) int {
	splitted := splitInt(y)
	result := 0

	for i := 0; i < len(splitted); i++ {
		actual := x * splitted[i]
		result += int(float64(actual) * math.Pow(10, float64(i)))
	}

	return result
}