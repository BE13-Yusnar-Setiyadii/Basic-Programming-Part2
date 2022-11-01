package main

import "fmt"

func FaktorBilanganDesc(n int) string {
	var faktor string
	for i := n; i > 0; i-- {
		if n%i == 0 {
			faktor += fmt.Sprintln(i)
		}
	}
	return faktor
}

func main() {
	var n int = 6
	fmt.Println(FaktorBilanganDesc(n))
}
