package main

import "fmt"

func largestPrimeFactor(number int) int {
	i := 2
	for number > 1 {
		if number % i == 0 {
			number /= i
		}else{
			i += 1
		}
	}
	return i
}
func main() {
	foo := 600851475143
	fmt.Printf("The largest prime factor of %d is %d", foo, largestPrimeFactor(foo))
}
