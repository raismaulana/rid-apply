package logic_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Reverse() {
	fmt.Println("Masukkan kalimat: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	sentence := scanner.Text()
	words := strings.Fields(sentence)
	var sdrow string
	//loop each word
	for _, word := range words {
		var capIndex = make([]int, 0)
		//save index capital letter and reverse word
		var backword string
		lowerWord := strings.ToLower(word)
		for i, v := range word {
			if unicode.IsUpper(v) {
				capIndex = append(capIndex, i)
			}
			backword = string(lowerWord[i]) + backword
		}
		//capitalize the letter that refers to saved index
		var capBackword string
		for i, v := range backword {
			if isIn(capIndex, i) {
				v = unicode.ToUpper(v)
			}
			capBackword = capBackword + string(v)

		}
		//concat word
		sdrow = sdrow + capBackword + " "
	}
	fmt.Println(sdrow)
}

func isIn(slice []int, i int) bool {
	for _, item := range slice {
		if item == i {
			return true
		}
	}
	return false
}
