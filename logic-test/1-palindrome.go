package logic_test

import (
	"fmt"
	"strings"
)

// Palindrome is a word, number, phrase, or other sequence of characters which reads the same backward as forward
func Palindrome(word string) {
	lowerWord := strings.ToLower(word)
	var backword string
	for _, v := range lowerWord {
		backword = string(v) + backword
	}

	if lowerWord == backword {
		fmt.Printf("%s adalah palindrom\n", word)
	} else {
		fmt.Printf("%s adalah bukan palindrom\n", word)
	}
}
