package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("starting Go learning Day#1")
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(multiply(10, 9))
	fmt.Println(mutliply_shorthand(10, 11))
	fmt.Println(multi_return(10, 11))

	var i int
	fmt.Println(i)

	k := 7
	fmt.Println("Dhoni jersey no: ", k)
}

func multiply(x int, y int) int {
	return x * y
}

func mutliply_shorthand(x, y int) int {
	return x * y
}

func multi_return(x, y int) (int, int) {
	return x * y, x + y
}
