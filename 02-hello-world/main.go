package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world 2")
	foo()
}

func foo() {
	fmt.Print("Foo is called \n")
	bar()
}

func bar() {
	fmt.Println("Bar is called")
}
