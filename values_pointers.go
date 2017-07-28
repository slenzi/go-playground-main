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

func pointerTest1(){

	a := 5		// standard variable of type int
	b := a		// another variable of type int
	c := &a		// a variable of type 'pointer to int'

	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n *c = %v\n", a, b, c, *c)
	fmt.Printf("Types:\n a:  %T\n b:  %T\n c: %T\n*c:  %T\n", a, b, c, *c)

	a = 7		// value of b will not change, but value of c will change because c points to the memory space of a

	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n *c = %v\n", a, b, c, *c)
	fmt.Printf("Types:\n a:  %T\n b:  %T\n c: %T\n*c:  %T\n", a, b, c, *c)

	// dereference c and indirectly change a
	*c = 4		// will change the value of a because c points to the memory space of a

	fmt.Printf("Values:\n a = %v\n b = %v\n c = %v\n *c = %v\n", a, b, c, *c)
	fmt.Printf("Types:\n a:  %T\n b:  %T\n c: %T\n*c:  %T\n", a, b, c, *c)

}

func pointerTest2(){

	p := Vertex{1, 2}  // has type Vertex
	q := &Vertex{1, 2} // has type *Vertex
	r := Vertex{X: 1}  // Y:0 is implicit
	s := Vertex{}      // X:0 and Y:0
	t := q
	v := *q
	q.X = 4
	u := *q


	fmt.Printf("p = %v, q = %v, r = %v, s = %v, t = %v, u = %v, (*t == u) = %v, v = %v, (v == u) = %v\n", p, q, r, s, t, u, *t == u, v, v == u)

}

func pointerTest3(){

	a := 2
	b := newInt1()
	c := newInt2()

	fmt.Printf("Values: a = %v, b = %v, c = %v\n", a ,b, c)
	fmt.Printf("Types: a = %T, b = %T, c = %T, *b = %T, *c = %T\n", a, b, c, *b, *c)

	*b = 3
	c = &a

	fmt.Printf("Values: a = %v, b = %v, c = %v\n", a ,*b, *c)
	fmt.Printf("Types: a = %T, b = %T, c = %T\n", a, *b, *c)


}
// example use of new
func newInt1() *int {
	return new(int)
}
func newInt2() *int {
	var i int
	return &i
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
