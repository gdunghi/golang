package golang

import (
	"reflect"
	"testing"
)

func TestOddList(t *testing.T) {
	o := odds(5.0)
	expected := []float64{1.0, 3.0, 5.0, 7.0, 9.0}

	if !reflect.DeepEqual(expected, o) {
		t.Errorf("%v is expected but got %v", expected, o)
	}
}

func TestOddList2(t *testing.T) {
	o := odds(7.0)
	expected := []float64{1.0, 3.0, 5.0, 7.0, 9.0, 11.0, 13.0}

	if !reflect.DeepEqual(expected, o) {
		t.Errorf("%v is expected but got %v", expected, o)
	}
}

func TestInvert(t *testing.T) {
	oe := newOddE()

	if oe.invert() != 1.0 {
		t.Error("it should be 1.0 at first time but got", oe)
		return
	}

	if oe.invert() != -1.0 {
		t.Error("it should be 1.0 at first time but got", oe)
		return
	}

	if oe.invert() != 1.0 {
		t.Error("it should be 1.0 at first time but got", oe)
		return
	}

	if oe.invert() != -1.0 {
		t.Error("it should be 1.0 at first time but got", oe)
		return
	}
}

// func TestPi(t *testing.T) {
// 	p := pi(odds(10000.0), newOddE())

// 	if p != 3.141392653591791 {
// 		t.Error("it should be 3.141392653591791 but got", p)
// 	}
// }
