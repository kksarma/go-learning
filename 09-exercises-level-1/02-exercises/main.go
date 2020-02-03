package main

import (
	"fmt"
)

func main() {
	exercises1()
	exercises2()
}

var x int = 42
var y string = "James Bond"
var z bool = true

func exercises1() {
	fmt.Println("************* Exerices 3 *************")
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

type krishan int

var a krishan
var b int

func exercises2() {
	fmt.Println("\n************* Exerices 4-5 *************")

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	a = 42
	b = int(a)
	fmt.Println(a)
	fmt.Println(b)
}
