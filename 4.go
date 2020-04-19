package main

import (
	"fmt"
	"strconv"
)

//isPalindrome takes an int as input
//returns a bool reflecting whether ot not the int
//is a palindrome
func isPalindrome(number int) string {
	numberString := strconv.Itoa(number)
	reversedString := []string{}
	for i := range numberString {
		i = len(numberString) - 1 - i
		reversedString = append(reversedString, numberString[i])
	}
	return numberString
}

func main() {
	foo := 1234
	stringFoo := isPalindrome(foo)
	fmt.Printf("The string version of foo is %q", stringFoo)
}
