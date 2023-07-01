package main

import (
	"fmt"
)

func main() {

	fizzBuzzTest()

	word := "Aibohphobia"
	if palindromeTest(word) {
		fmt.Println(word, "adalah palindrome")
	} else {
		fmt.Println(word, "bukan palindrome")
	}
}
