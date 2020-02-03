package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World 3")
	evenNumber()
}

func evenNumber() {
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
