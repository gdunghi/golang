package main

import "strconv"

func fizzBuzz(n int) string {
	if isFizz(n) && isBuzz(n) {
		return "FizzBuzz"
	} else if isFizz(n) {
		return "Fizz"
	} else if isBuzz(n) {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func isFizz(n int) bool {
	if n%3 == 0 {
		return true
	}

	return false
}

func isBuzz(n int) bool {
	if n%5 == 0 {
		return true
	}

	return false
}
