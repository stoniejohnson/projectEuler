package main

import (
	"fmt"
	"math"
)

//sumOfSquares calculates the sum of the squares from 1 
//up to and including the number provided
func sumOfSquares(number float64) float64 {
	sum := 0.0 //instantiate initial sum
	for i := 1.0; i <= number; i++ {
		square := math.Pow(i,2)
		sum += square
	}
	return sum
}		
//squareOfSum calculates the sum of the numbers up to
//and including the number provided, and then returns
//its square
func squareOfSum(number float64) float64 {
	sum := 0.0 //instantiate initial sum
	for i := 1.0; i <= number; i++ {
		sum += i
	}
	return math.Pow(sum,2)
}

func main() {
	fmt.Printf("The answer is %d\n", int(squareOfSum(100)-sumOfSquares(100)))
}
