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
	largest := 0
	for i := 900; i < 1000; i++ {
		for j := 900; j < 1000; j++ {
			if isPalindrome(i * j) && (i * j) > largest {
				largest = i * j
			}
		}
	}
	fmt.Printf("The largest palindrome is %d\n", largest)
}
