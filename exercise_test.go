package golang

import "testing"

func TestFactorial1(t *testing.T) {
	result := factorial(3)
	if result != 6 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial2(t *testing.T) {
	result := factorial(4)
	if result != 24 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial3(t *testing.T) {
	result := factorial(5)
	if result != 120 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial4(t *testing.T) {
	result := factorial(6)
	if result != 720 {
		t.Error("Please make the factorial function")
	}
}
