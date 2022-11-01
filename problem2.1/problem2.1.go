package main

import "fmt"

func FaktorBilangan(n int) string {
	var faktor string
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor += fmt.Sprintln(i)
		}
	}
	return faktor
}

func main() {
	var n int = 6
	fmt.Println(FaktorBilangan(n))
}
