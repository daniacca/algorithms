package integer_multiplication

import "math"

func RecursiveMultiply(x, y int) int {
	if x < 10 || y < 10 {
		return x * y
	}

	n := int(math.Max(float64(len(splitInt(x))), float64(len(splitInt(y)))))
	m := int(math.Ceil(float64(n) / 2))

	a, b := SplitAt(float64(x), m)
	c, d := SplitAt(float64(y), m)

	ac := RecursiveMultiply(a, c)
	ad := RecursiveMultiply(a, d)
	bc := RecursiveMultiply(b, c)
	bd := RecursiveMultiply(b, d)
	
	return int(float64(ac) * math.Pow(10, float64(2*m)) + float64(ad + bc) * math.Pow(10, float64(m)) + float64(bd))
}