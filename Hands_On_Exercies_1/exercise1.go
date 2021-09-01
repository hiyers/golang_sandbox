package main

import "fmt"

func main() {
	var x int = 42
	var y string = "James Bond"
	var z bool = true

	fmt.Println(x, y, z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
