package main

import (
	"fmt"
//	"math"
)

func isPrime(number int) bool {
	if number == 1 {
		return false
	}
	
	for i := 2; i < number; i ++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	testNumber := 2
	if isPrime(testNumber) {
		fmt.Printf("%d is prime\n", testNumber)
	}
	if !isPrime(testNumber) {
		fmt.Printf("%d is not prime\n", testNumber)
	}
}
 
 
