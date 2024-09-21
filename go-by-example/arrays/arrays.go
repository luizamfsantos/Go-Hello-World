package main

import "fmt"

func main() {
	var a [5]int
	fmt.Printf("empty: %v\n", a)

	a[4] = 100
	fmt.Printf("set: %v\n", a)
	fmt.Printf("get: %v\n", a[4])

	fmt.Printf("lenght: %v\n", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // if you specify the index with :,
	// the elements in between will be zeroed
	fmt.Println("idx:", b)

	twoD := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("two dimensional array %v\n", twoD)
	fmt.Println("The second element of the first row is ", twoD[0][1])
	fmt.Println("The third element of the 2nd row is ", twoD[1][2])
}
