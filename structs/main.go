package main

import "fmt"

/*
use pointer to change these in function [int, arr, float, bool, string, strcuts]
no need to used pointers [slices, maps, channels, pointers, functions]
*/

type person struct {
	firstname string
	lastname  string
	contact contactInfo
}

type contactInfo struct {
	pin   int
	email string
}

func (p person) print() {
	fmt.Printf("%+v\n\n", p)
}

func (p *person) update(name string) {
	// (*p).firstname = name
	p.firstname = name //shorthand dot operator on struct automatically dereferences
}
func main() {
	var p1 person
	p2 := person{"shubham", "dogra", contactInfo{1234, "abc@gmail.com"}}
	p3 := person{
		firstname: "Yuvraj",
		lastname:  "Sharma",
		contact: contactInfo{
			pin:   123456,
			email: "abc@gmail.com",
		},
	}

	//printing
	p1.print()
	p2.print()
	p3.print()
	/*
		updating
		structs are always passed by value
		means when we pass struct to function it will create a copy of that
		in a diff mem location and updated are not reflected where
		we want them to
		therfore we use pointer
	*/
	p1add := &p1 //pointer of type person
	p1add.update("Akshat")

	//shortcut
	p2.update("Anubhav")
	p1.print()
	p2.print()
}
