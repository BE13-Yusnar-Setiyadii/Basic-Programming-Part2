package main

import "fmt"

func KonversiNilai(nilai int) string {
	if nilai >= 80 && nilai <= 100 {
		return "A"
	} else if nilai >= 65 && nilai < 80 {
		return "B+"
	} else if nilai >= 50 && nilai < 65 {
		return "B"
	} else if nilai >= 35 && nilai < 50 {
		return "C"
	} else if nilai >= 0 && nilai < 35 {
		return "D"
	} else {
		return "tidak ada nilai"
	}
}

func main() {
	var nilai int = 80

	fmt.Println(KonversiNilai(nilai))
}
