package main

import "fmt"

func isPalindrome(x int) bool {
	// if x is negative, it cannot be a palindrome
	if x < 0 {
		return false
	}

	// write x as a slice of digits
	digits := []int{}
	for i := 0; i < 10; i++ {
		if x < 10 {
			digits = append(digits, x)
			break
		}
		digits = append(digits, x%10)
		x = x / 10
	}

	// check if the digits are a palindrome
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	answer1 := isPalindrome(121)
	fmt.Printf("Answer: %t\n", answer1) // true

	answer2 := isPalindrome(-121)
	fmt.Printf("Answer: %t\n", answer2) // false

	answer3 := isPalindrome(10)
	fmt.Printf("Answer: %t\n", answer3) // false

	answer4 := isPalindrome(321)
	fmt.Printf("Answer: %t\n", answer4) // false
}
