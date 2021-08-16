package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		indexClosingBracketFound := strings.Index(str, ")")

		if indexFirstBracketFound >= 0 && indexClosingBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound+1 : indexClosingBracketFound])

			return wordsAfterFirstBracket
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func main() {
	result := findFirstStringInBracket("Kamu makannya apa? (SATE!!)")
	fmt.Println(result)
}
