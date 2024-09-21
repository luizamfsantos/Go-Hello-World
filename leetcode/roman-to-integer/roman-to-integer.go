package main

import "fmt"

var roman_mapping = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var total int
	for i := 0; i < len(s); i++ {
		string_subset := string(s[i])
		val := roman_mapping[string_subset]
		total += val
	}

	return total
}

func main() {
	string1 := "III"
	answer1 := romanToInt(string1)
	fmt.Printf("%s is %d\n", string1, answer1) // III is 3

	string2 := "LVIII"
	answer2 := romanToInt(string2)
	fmt.Printf("%s is %d\n", string2, answer2) // LVIII is 58

	string3 := "MCMXCIV"
	answer3 := romanToInt(string3)
	fmt.Printf("%s is %d\n", string3, answer3) // MCMXCIV is 1994
}
