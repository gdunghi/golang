package golang

import (
	"strings"
)

// LettersCount is function
func LettersCount(sentenses string) map[string]int {
	var m = map[string]int{}
	split := strings.Split(sentenses, " ")

	for _, v := range split {
		m[v] = len(v)
	}

	return m
}
