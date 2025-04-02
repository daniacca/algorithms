package integer_multiplication

import "testing"

func TestKaratsubaMultiply(t *testing.T) {
	x := 1234
	y := 5678
	expected := 7006652
	actual := KaratsubaMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestKaratsubaMultiplyZero(t *testing.T) {
	x := 1234
	y := 0
	expected := 0
	actual := KaratsubaMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestKaratsubaMultiplyBig(t *testing.T) {
	x := 12345678
	y := 87654321
	expected := 1082152022374638
	actual := KaratsubaMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func BenchmarkKaratsubaMultiply(b *testing.B) {
	x := 12345678
	y := 87654321
	for i := 0; i < b.N; i++ {
		KaratsubaMultiply(x, y)
	}
}