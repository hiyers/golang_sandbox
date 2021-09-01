package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("5 + 6 =", mySum(5, 6))
	fmt.Println("6 + 9 =", mySum(6, 9))

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum

}
