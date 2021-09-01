package main

import "fmt"

type age int

var x age

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)

}
