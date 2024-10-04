package main

import "fmt"

// anonymous functions
// used whenever we want to define a function
// without hanging to name it

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // different from nextInt
	fmt.Println(newInts())
}
