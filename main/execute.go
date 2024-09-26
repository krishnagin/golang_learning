package main

import (
	"fmt"
	"math"
	"strings"
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

	//Data types
	var b bool = true
	fmt.Println(b)

	run := 50
	float_run := float32(run)
	fmt.Println(float_run)

	v := true                          // change me!
	fmt.Printf("v is of type %T\n", v) // %T appears to be special character which prints the type of the variable

	const DAYS = 7

	sum := 0

	for i := 0; i < 1000; i++ {
		sum += i
	}
	fmt.Println(sum)

	if score := 98; score > 50 && score < 100 {
		fmt.Println("fifty scored")
	}

	sumP := &sum
	fmt.Println(sumP)

	type Over struct {
		bowler string
		runs   int
	}
	fmt.Println(Over{}) // Empty initialization of data type Over
	fmt.Println(Over{"starc", 7})
	o := Over{"starc", 7}
	fmt.Println(o.bowler)
	fmt.Println(o.runs) // refer with . the properties of the struct data type

	oP := &o
	fmt.Println(oP) // this is not printing an address unlike of built-in data types
	fmt.Println(oP.bowler)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	q := []int{2, 3, 5, 7, 11, 13} // if no number is specified, then it's slice
	fmt.Println(q)

	m := make([]int, 5)
	fmt.Printf("%v %v", len(m), cap(m))

	m = m[1:]
	fmt.Println()
	fmt.Printf("%v %v", len(m), cap(m))
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

func WordCount(s string) map[string]int {
	sA := strings.Fields(s)
	fmt.Println(sA)
	result := make(map[string]int)

	for _, v := range sA {
		_, ok := result[v]
		if ok {
			result[v] = result[v] + 1
		} else {
			result[v] = 1
		}
	}
	return result
}
