package logic_test

import (
	"fmt"
	"strings"
)

// Palindrome is a word, number, phrase, or other sequence of characters which reads the same backward as forward
func Palindrome(word string) bool {
	lowerWord := strings.ToLower(word)
	var backword string
	for _, v := range lowerWord {
		backword = string(v) + backword
	}

	if lowerWord == backword {
		fmt.Printf("%s is palindrom\n", word)
		return true
	} else {
		fmt.Printf("%s is not palindrom\n", word)
		return false
	}
}
