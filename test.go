package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {

	fmt.Printf("hello, world! This is golang.\n")
	fmt.Println("Another line of text....")
	fmt.Println("The time is ->", time.Now())
	fmt.Println("The time in milli is ->", makeTimestamp())
	fmt.Printf("The number %g roughly represents Pi!\n",math.Pi)

	fmt.Printf("The numbers %v + %v = %v", 3, 7, add(3,7))

	rand.Seed(makeTimestamp())
	fmt.Println("My favorite number is", rand.Intn(10))

	var x1 string = "hello"
	var x2 string = "world"

	a, b := swap(x1, x2)
	fmt.Printf("Swap (%v, %v) gives you (%v, %v)\n",x1,x2,a,b)

	loopTests()
	ternaryTest()
	switchTest()
	sliceTest()
	valuesPointersTests()

	val, err := deferWithNamedReturns(11)

	fmt.Printf("Returned value = %v\n", val)
	if(err != nil) {
		fmt.Printf("Returned error = %v\n", err)
	}

	deferStackDemo()

	methodsTests()

}

// comment
func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

/**
 Multi line comment
 Return the addition of two numbers
 */
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func loopTests(){

	fmt.Println("Begin loop tests.")

	fmt.Println("Print even integers from 0 to 10.")
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("End loop tests.")

}

// golang has no ternary
func ternaryTest(){

	fmt.Println("Begin ternary tests.")

	//var c = (a > b) ? a : b;
	a := 5
	b := 10

	c := a
	if a > b {
		c = a
	}

	fmt.Println(c)

	fmt.Println("End ternary tests.")

}

func switchTest(){

	fmt.Println("Begin switch tests.")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case float32:
			fmt.Println("I'm a float32")
		case float64:
			fmt.Println("I'm a float64")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	var f1 float32 = 1.3
	var f2 float64 = 1.7
	whatAmI(true)
	whatAmI(1)
	whatAmI(f1)
	whatAmI(f2)
	whatAmI("hey")

	fmt.Println("End switch tests.")

}

func sliceTest(){

	fmt.Println("Begin slice tests")

	mySlice := make([]string, 3)
	fmt.Println("emp:", mySlice)

	mySlice[0] = "a"
	mySlice[1] = "b"
	mySlice[2] = "c"
	mySlice = append(mySlice, "d")
	mySlice = append(mySlice, "e", "f")
	fmt.Println("apd:", mySlice)

	for index, value := range mySlice {
		fmt.Printf("Value at index %d = %v\n", index, value)
	}

	c := make([]string, len(mySlice))
	copy(c, mySlice)
	fmt.Println("cpy:", c)

	l := mySlice[2:5]
	fmt.Println("sl1:", l)

	fmt.Println("End slice tests")

}

func valuesPointersTests(){

	fmt.Println("Begin values & pointers tests")

	pointerTest1()
	pointerTest2()
	pointerTest3()

	passByValueTest1()
	passByReferenceTest1()
	passByValueTest2()
	passByReferenceTest2()

	fmt.Println("End values & pointers tests")

}

func deferWithNamedReturns(x int) (id int, err error) {
	defer func() {
		if id == 10 {
			err = fmt.Errorf("Invalid Id\n")
		}else{
			fmt.Printf("Provided number %v is good\n", x)
		}
	}()

	id = x

	return
}

// defer calls are pushed on a stack, so calls are executed in last-in-first-out order
func deferStackDemo(){

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

func methodsTests(){

	v1 := getSampleVertexValue()
	v2 := getSampleVertexPointer()

	fmt.Printf("Vertex = %+v\n", v1)
	fmt.Printf("Vertex = %+v\n", v2)

}