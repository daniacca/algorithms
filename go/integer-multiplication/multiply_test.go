package integer_multiplication

import "testing"

func TestMultiply(t *testing.T) {
	x := 1234
	y := 5678
	expected := int64(7006652)
	actual := Multiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestMultiplyZero(t *testing.T) {
	x := 1234
	y := 0
	expected := int64(0)
	actual := Multiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestMultiplyBig(t *testing.T) {
	x := 123456789
	y := 987654321
	expected := int64(121932631112635269)
	actual := Multiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func BenchmarkMultiply(b *testing.B) {
	x := 123456789
	y := 987654321
	for i := 0; i < b.N; i++ {
		Multiply(x, y)
	}
}
