package golang

import (
	"reflect"
	"testing"
)

func TestFibonacci1(t *testing.T) {
	result := fibonacci(5)

	expected := []int{0, 1, 1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Please make the fibonacci function")
		t.Errorf("%v\nis my expected but got\n%v\n", expected, result)
	}
}

func TestFibonacci2(t *testing.T) {
	result := fibonacci(10)

	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Please make the fibonacci function")
		t.Errorf("%v\nis my expected but got\n%v\n", expected, result)
	}
}
