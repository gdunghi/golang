package golang

import (
	"errors"
	"fmt"
)

func echo(s fmt.Stringer) {
	print(s.String())
}

func throw(e error) {
	print(e.Error())
}

type iface interface {
	Name() string
	Age() int
}

func needIF(i iface) {
	fmt.Println(i.Name(), i.Age())
}

type myStringer interface {
	fmt.Stringer
	error
}

func strErr(s myStringer) {
	fmt.Println(s.Error(), s.String())
}

type myInteface interface {
	String() string
	get() error
}

type err struct {
	Err error
}

//String aa
func (p person) String() string {
	return p.Name
}

func (p person) get() error {
	return errors.New(p.Name)
}

type person struct {
	Name string
}

func newStringer(name string) fmt.Stringer {
	p := person{Name: name}
	return p
}

func newThrow(str string) person {
	p := person{Name: str}
	return p
}
