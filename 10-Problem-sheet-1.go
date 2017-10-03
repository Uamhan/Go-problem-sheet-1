package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "hello how are you today." // stringto reverse

	//prints the string andthe reverseof the string using the reverseStringFunction
	fmt.Printf("original string : %v \n", word)
	fmt.Printf("reversed string : %v \n", reverseString(word))
}

func reverseString(word string) string {

	wordLen := utf8.RuneCountInString(word) //gets length of string
	runes := make([]rune, wordLen)          // changes the string to a slice of runes

	for _, rune := range word { //iterates though the string
		wordLen--             //decremets word length
		runes[wordLen] = rune //swaps the current rune with runes[wordLen]
	}

	return string(runes[wordLen:]) //returns the reversed string

}
