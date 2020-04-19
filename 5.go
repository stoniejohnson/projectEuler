package main

import "fmt"

func isDivisible(number int) bool {
	for i := 1; i <= 20; i++ {
		if number % i != 0 {
			return false
		}
	}		
	return true

}

func main() {
	number := 2520
	notFound := true
	for notFound {
		if isDivisible(number) {
			fmt.Printf("You lucky duck, ya found it: %d\n", number)
			return
		}
		fmt.Printf("Yeah man, %d just didn't cut it. Let's try %d\n", number, number + 1)
		number += 2520
	}		
}
