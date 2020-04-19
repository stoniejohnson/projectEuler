package main

import "fmt"

func main() {
	foo := fibEvenSum(4000000)
	fmt.Printf("The sum is %d", foo)
}

func fibEvenSum(length int) int {
	// initialize values
	a := 1
	b := 2

	// instantiate sum
	sum := 2
	for a+b < length {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
			fmt.Printf("Just added %d to the sum, the sum is %d\n", b, sum)
		}
	}

	return sum
}
