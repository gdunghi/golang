package golang

import "testing"

func TestStringerEcho(t *testing.T) {
	s := newStringer("my name")
	echo(s)
}

func TestThrowError(t *testing.T) {
	e := newThrow("this is an error")
	err := e.get()
	throw(err)
}
