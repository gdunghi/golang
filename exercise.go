package golang

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
