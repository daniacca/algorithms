package integer_multiplication

import "math"

func KaratsubaMultiply(x, y int) int {
	if x < 10 || y < 10 {
		return x * y
	}

	n := int(math.Max(float64(len(splitInt(x))), float64(len(splitInt(y)))))
	m := RoundUp(float64(n) / 2)

	a, b := SplitAt(float64(x), m)
	c, d := SplitAt(float64(y), m)

	p := a + b
	q := c + d

	ac := KaratsubaMultiply(a, c)
	bd := KaratsubaMultiply(b, d)
	pq := KaratsubaMultiply(p, q)
	adbc := pq - ac - bd

	return int(float64(ac)*math.Pow(10, float64(2*m)) + float64(adbc)*math.Pow(10, float64(m)) + float64(bd))
}