package integer_multiplication

import "math"

// Split a number into two numbers at a given index
func SplitAt(num float64, ind int) (int, int) {
	a := int(num / math.Pow(10, float64(ind)))
	b := int(num) - a*int(math.Pow(10, float64(ind)))
	return a, b
}

// Rounds up to the nearest whole number
func RoundUp(num float64) int {
	return int(math.Ceil(num))
}