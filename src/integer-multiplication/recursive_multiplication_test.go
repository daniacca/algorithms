package integer_multiplication

import "testing"

func TestRecursiveMultiply(t *testing.T) {
	x := 1234
	y := 5678
	expected := 7006652
	actual := RecursiveMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestRecursiveMultiplyZero(t *testing.T) {	
	x := 1234
	y := 0
	expected := 0
	actual := RecursiveMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestRecursiveMultiplyBig(t *testing.T) {
	x := 12345678
	y := 87654321
	expected := 1082152022374638
	actual := RecursiveMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func BenchmarkRecursiveMultiply(b *testing.B) {
	x := 12345678
	y := 87654321
	for i := 0; i < b.N; i++ {
		RecursiveMultiply(x, y)
	}
}