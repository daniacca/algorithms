package integer_multiplication

import "testing"

func TestShoolAlgo(t *testing.T) {
	x := 1234
	y := 5678
	expected := 7006652
	actual := SchoolMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestShoolAlgoZero(t *testing.T) {
	x := 1234
	y := 0
	expected := 0
	actual := SchoolMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestShoolAlgoBig(t *testing.T) {
	x := 123456789
	y := 987654321
	expected := 121932631112635269
	actual := SchoolMultiply(x, y)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func BenchmarkShoolAlgo(b *testing.B) {
	x := 123456789
	y := 987654321
	for i := 0; i < b.N; i++ {
		SchoolMultiply(x, y)
	}
}