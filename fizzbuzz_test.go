package main

import (
	"testing"
)

func TestFizz(t *testing.T) {
	expected := "Fizz"
	result := fizzBuzz(3)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestBuzz(t *testing.T) {
	expected := "Buzz"
	result := fizzBuzz(5)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestFizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	result := fizzBuzz(15)
	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestNotFizzBuzz(t *testing.T) {
	expected := "7"
	result := fizzBuzz(7)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}
