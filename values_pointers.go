package main

import "fmt"

/*

Allowed:
--------
p := Person{"Steve", 28} 	stores the value
p := &Person{"Steve", 28} 	stores the pointer address (reference)
PrintPerson(p) 			passes either the value or pointer address (reference)
PrintPerson(*p) 		passes the value
PrintPerson(&p) 		passes the pointer address (reference)
func PrintPerson(p Person)	ONLY receives the value
func PrintPerson(p *Person)	ONLY receives the pointer address (reference)

Not Allowed:
--------
p := *Person{"Steve", 28} 	illegal
func PrintPerson(p &Person)	illegal

*/

type Person struct {
	Name string
	Age int
}

// *****************************************************************************
// Pass by Value 2
// *****************************************************************************

func passByValueTest1() {
	fmt.Println("Pass by value test 1")
	p := Person{"Steve", 28}
	printPerson1(p)
	updatePerson1(p)
	printPerson1(p)
}

func updatePerson1(p Person) {
	p.Age = 32
	printPerson1(p)
}

func printPerson1(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Pass by Reference 1
// *****************************************************************************

func passByReferenceTest1() {
	fmt.Println("Pass by reference test 1")
	p := &Person{"Steve", 28}
	printPerson2(p)
	updatePerson2(p)
	printPerson2(p)
}

func updatePerson2(p *Person) {
	p.Age = 32
	printPerson2(p)
}

func printPerson2(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Pass by Value 2 (requires more typing)
// *****************************************************************************

func passByValueTest2() {
	fmt.Println("Pass by value test 2")
	p := &Person{"Steve", 28}
	printPerson4(*p)
	updatePerson4(*p)
	printPerson4(*p)
}

func updatePerson4(p Person) {
	p.Age = 32
	printPerson4(p)
}

func printPerson4(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Pass by Reference 2 (requires more typing)
// *****************************************************************************

func passByReferenceTest2() {
	fmt.Println("Pass by reference test 2")
	p := Person{"Steve", 28}
	printPerson3(&p)
	updatePerson3(&p)
	printPerson3(&p)
}

func updatePerson3(p *Person) {
	p.Age = 32
	printPerson3(p)
}

func printPerson3(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}
