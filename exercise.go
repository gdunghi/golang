package golang

// formula: π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...

func odds(f float64) []float64 {
	var result = []float64{}

	sum := 0
	for len(result) < int(f) {
		if sum%2 != 0 {
			result = append(result, float64(sum))

		}
		sum++

	}

	return result
}

func newOddE() odde {

}

func (o odde) invert(float64) float64 {
	return o.Number
}

type odde struct {
	Number float64
}
