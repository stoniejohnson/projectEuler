package main

import (
	"fmt"
	"strconv"
)

//isPalindrome takes an int as input
//returns a bool reflecting whether ot not the int
//is a palindrome
func isPalindrome(number int) bool {
	numberString := strconv.Itoa(number)
	reversedString := []byte{}
	for i := range numberString {
		i = len(numberString) - 1 - i
		reversedString = append(reversedString, numberString[i])
	}
	if numberString == string(reversedString) {
		return true 
	}else{
		return false
	}
}

func main() {
	foo := 1234
	fmt.Printf("Is %d a palindrome?\n %t\n", foo,  isPalindrome(foo))
	yeet := 1111
	fmt.Printf("Is %d a palindrome?\n %t\n", yeet, isPalindrome(yeet))
}
