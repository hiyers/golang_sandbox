package main

import (
	"fmt"
	"runtime"
)

type age int

var x age

var y int

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	y = int(x)
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
