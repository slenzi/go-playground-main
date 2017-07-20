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