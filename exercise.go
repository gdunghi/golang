package golang

func enc(x int, n int) int {
	if x == 32 {
		return x
	}

	return (x + n) % 26

}

func caesar(s string, rot int) string {
	result := ""
	for _, v := range s {
		if v == 32 {
			result += string(32)
		} else {
			m := ((int(v) + rot) % 26)
			result += string(m + 97)
		}
	}
	return result
}
