package main

import (
	"fmt"
)

func main() {
	// We can't keep unsed variables in GO otherwise GO compiler throws error `xyz` declared and not used.
	// We use _ operator for throw away funtions return values

	n, _ := fmt.Println("Hello world", 67, 7, true, 89.079)
	fmt.Println(n)

	// Here := is a short declaration operator in which declaration and assignment done same time.
	y := 23
	fmt.Println(y)

	x := 10 + y
	fmt.Println(x)

	z := x + y
	fmt.Println(z)

	s := "GoLang " + "Hi"
	fmt.Println(s)
}
