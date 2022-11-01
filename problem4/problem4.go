package main

import "fmt"

func Palindrome(input string) bool {
	str1 := ""
	for i := len(input) - 1; i >= 0; i-- {
		str1 += string(input[i])
	}
	for i := range input {
		if input[i] != str1[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Palindrome("cavic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
