package main

import (
	"fmt"
)

func main() {
	exercises1()
	exercises2()
}

func exercises1() {
	fmt.Println("************* Exerices 1 *************")
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

var x int
var y string
var z bool
func exercises2() {
	fmt.Println("\n************* Exerices 2 *************")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}