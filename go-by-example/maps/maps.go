package main

import (
	"fmt"
	"maps"
)

// Go's version of a hash table/dict

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m) // printing a map shows all k/v pairs
	fmt.Printf("Location of m: %p\n", &m)

	v1, prs1 := m["k1"] // get value similar to any hash
	fmt.Println("v1: ", v1)
	fmt.Printf("Location of v1: %p\n", &v1)
	fmt.Println("prs1: ", prs1)

	v3 := m["k3"] // if key doesn't exist, it returns zero value
	fmt.Println("v3: ", v3)

	fmt.Println("len: ", len(m)) // len returns the number of pairs

	delete(m, "k2")                       // remove the k/v
	fmt.Printf("Location of m: %p\n", &m) // delete operation is in-place
	fmt.Println("map: ", m)

	clear(m) // removes all k/v pairs
	fmt.Println("map: ", m)

	_, prs := m["k2"]         // the optional second return is indicator (present)
	fmt.Println("psr: ", prs) // bool

	n := map[string]int{"foo": 1, "bar": 2} // declare and initialize
	fmt.Println("map: ", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}

}
