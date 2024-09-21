package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500_000_000
	
	const d = 3e20 / n // Constant expressions perform arithmetic with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) 	// a numeric constant has no type until it's given one, such as
							// by an explicit conversion 

	fmt.Println(math.Sin(n)) // a number can be given a type by using it in a context that 
							 //requires one, such as a variable assignment or function call
}