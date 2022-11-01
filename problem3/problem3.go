package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	for i := 2; i <= number-1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(PrimeNumber(35))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
