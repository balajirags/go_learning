package main

import (
	"os"
	"fmt"
)

func main() {
	input := os.Args[1]
	fmt.Printf("Is the word %s a palindrome? - %t", input, isPalindorme(input))
}

func isPalindorme(input string) bool {
	chars := []rune(input)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %d, j=%d \n", i, j)
		fmt.Printf("i = %c, j = %c \n", chars[i], chars[j])
		if chars[i] != chars[j]{
			return false;
		}
	}
	return true
}
