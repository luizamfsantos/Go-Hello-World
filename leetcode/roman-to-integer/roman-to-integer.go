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
	var total, new_val, previous_val int
	for i := 0; i < len(s); i++ {
		string_subset := string(s[i])
		switch i {
		case 0:
			new_val = roman_mapping[string_subset]
			total += new_val
		default:
			previous_val = new_val
			new_val = roman_mapping[string_subset]
			if previous_val < new_val {
				total += new_val
				total -= 2 * previous_val
			} else {
				total += new_val
			}
		}
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

	string4 := "IV"
	answer4 := romanToInt(string4)
	fmt.Printf("%s is %d\n", string4, answer4) // IV is 4

}
