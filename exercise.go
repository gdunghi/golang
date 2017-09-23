package golang

func fibonacci(a int) []int {
	//{\displaystyle F_{n}=F_{n-1}+F_{n-2},}
	result := []int{}

	for i := 0; i < a; i++ {
		if i == 0 || i == 1 {
			result = append(result, i)
		} else {
			result = append(result, result[i-1]+result[i-2])
		}
	}

	return result
}

func concat(sorce, add []string) []string {

	return append(sorce, add...)
}

func delFirst(a []string) []string {
	return a[1:]
}

func delLast(a []string) []string {
	return a[:len(a)-1]
}

func delSecond(a []string) []string {
	first := make([]string, len(a))
	first[0] = a[0]
	last := a[2:]

	return append(first, last...)
}
