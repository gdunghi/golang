package golang

func stepup(a, b int) (int, int) {
	return b, a + b
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}
