package main

import "fmt"

func main() {
	// var declares 1 or more variables at once
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d bool = true
	fmt.Println(d)
	// variables declared without a initialization are zero-values
	var e int
	fmt.Println(e) // 0

	f := "apple" // shorthand for declaring and initializing
	fmt.Println(f)
}