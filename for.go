package main

import "fmt"

func main() {

	// while loops
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i ++
	}

	// for loops range
	for j := 0; j < 3; j ++ {
		fmt.Println(j)
	}

	// alternative for Go 1.22 or higher
	//for i := range 3 {
	//	fmt.Println("range", i)
	//}
	
	// while loops without a condition
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}