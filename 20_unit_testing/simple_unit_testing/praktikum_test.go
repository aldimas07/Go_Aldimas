package simpleunittesting

import (
	"testing"
)

func TestAdditon(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Additon(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubstract(t *testing.T) {
	var expected int = 4
	var actual int = Substract(6,2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	var expected int = 24
	var actual int = Multiply(6,4)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}

func TestDivision(t *testing.T) {
	var expected int = 2
	var actual int = Division(8,4)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)
	}
}

