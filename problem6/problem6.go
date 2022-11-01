package main

import (
	"fmt"
)

func prima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FullPrima(n int) bool {
	for n > 0 {
		if !prima(n % 10) {
			return false
		}
		n = n / 10
	}
	return true
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
