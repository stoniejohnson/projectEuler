package main

import "fmt"

func main() {
	entry := []string{"Jack","John","Jones"}
	for i, val := range entry {
	  fmt.Printf("At position %d, the character %s is present\n", i, val)
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
}
