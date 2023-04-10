package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

// any type that has the methods specified byan interface is also of the interface type
// an interface says

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	//in go a Value can be of more than one TYPE
	// in this example human p1 is both human and person
	var x human
	x = p1
	fmt.Printf("%T\n", x)
}
