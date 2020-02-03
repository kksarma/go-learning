package main

import ("fmt")

var a int

/* Here we created own type by name `krishan` underlying of int */
type krishan int
var b krishan

func main() {
	/* Here both a and b variable holds integer value but we can't assign a's value to b because both have different type */
	a = 10
	b = 20

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Type conversion
	a = int(b)
	fmt.Printf("%T\n", a)
}