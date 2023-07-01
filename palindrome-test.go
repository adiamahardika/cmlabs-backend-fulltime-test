package main

import "strings"

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindromeTest(str string) bool {
	str = strings.ReplaceAll(strings.ToLower(str), " ", "")
	reversedStr := reverseString(str)
	return str == reversedStr
}
