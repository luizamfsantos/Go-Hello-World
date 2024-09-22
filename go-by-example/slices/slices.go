package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices are types by the elements they contain
	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	// to create an empty slice with non-zero length
	// use builtin make
	s = make([]string, 3)
	fmt.Println("emp:", s, s == nil, ",length:", len(s), ",capacity:", cap(s))

	// we can set and get just like arrays
	s[0] = "a"
	fmt.Println("len:", len(s))
	s[1] = "b"
	fmt.Println("len:", len(s))
	s[2] = "c"
	fmt.Println("len:", len(s))
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// s[3] = "d" // gives an error: index out of range [3] with length 3
	fmt.Printf("Location of s: %p\n", &s)
	s = append(s, "d") //appends returns a new bigger slice
	s = append(s, "e", "d")
	fmt.Println("apd:", s)
	fmt.Println("length:", len(s), ",capacity:", cap(s))
	fmt.Printf("Location of s: %p\n", &s) // same location

	c := make([]string, len(s)) // empty slice of length the same as slice s
	copy(c, s)                  // copy s into c
	fmt.Println("cpy:", c)

	l := s[:3] // slice of a slice
	fmt.Println("sl1:", l)
	fmt.Printf("Location of l: %p\n", &l) // different location
	fmt.Printf("")

	t := []string{"g", "h", "i"} // declare and initialize a var for a slice
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // 2 dimensional slice
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // length of inner slice can vary
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
