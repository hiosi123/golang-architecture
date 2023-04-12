package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("from a person - this is my name", sa.first)
}

// any type that has the methods specified byan interface is also of the interface type
// an interface says

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	fmt.Printf("%T\n", p1)

	//in go a Value can be of more than one TYPE
	// in this example human p1 is both human and person
	var x, y human

	x = p1
	y = sa1
	x.speak()
	y.speak()
	fmt.Println("-------")
	foo(x)
	foo(y)
	foo(p1)
	foo(sa1)
}
