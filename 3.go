package main

import "fmt"

func isPrime(number int) bool {
	if number == 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	result := isPrime(5)
	fmt.Printf("%t", result)
}
