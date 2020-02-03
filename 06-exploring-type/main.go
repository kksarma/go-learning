package main

import (
	"fmt"
)

// GO lang is static type language so we have to declare the type of variable so it can only hold this type value.
// It is not dynamic type language like JavaScript

var a = "I love Golang"
var b = 22

func main() {
	// Println formats using the default formats for its operands and writes to standard output.
	// Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.
	fmt.Println("Value: ", a)

	// Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.
	fmt.Printf("Type: %T\n", a)

	// Print formats using the default formats for its operands and writes to standard output.
	// Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.
	fmt.Print(a, "Print")

	fmt.Println("Value: ", b)
	fmt.Printf("Type: %T\n", b)

	c := true
	fmt.Println("Value: ", c)
	fmt.Printf("Type: %T\n", c)

	d := `Hi 
	"What the hell is going on?" 
	Hey`
	fmt.Println("Value: ", d)
	fmt.Printf("Type: %T\n", d)

	e := 3999.456999
	fmt.Println("Value: ", e)
	fmt.Printf("Type: %T\n", e)
}
