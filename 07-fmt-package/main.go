package main

import (
	"fmt"
)

var a int = 98
var b bool = true
var c string = "A"

func main() {
	fmt.Println("Value ", a)

	/* General formats */
	fmt.Printf("Type: %T\n", a)
	fmt.Printf("%v \n", a)
	fmt.Printf("%#v \n", a)
	//fmt.Printf("%%\n", b)

	/* Boolean format */
	fmt.Printf("\n%t\n", b)

	/* Integer formats */
	fmt.Printf("%b\n", a) // base 2
	fmt.Printf("%c\n", a) // the character represented by the corresponding Unicode code point
	fmt.Printf("%d\n", a) // base 10
	fmt.Printf("%o\n", a) // base 8
	fmt.Printf("%x\n", a) // base 16, with lower-case letters for a-f
	fmt.Printf("%X\n", a) // base 16, with upper-case letters for A-F

	name, age := "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")
	fmt.Print(s)
}
