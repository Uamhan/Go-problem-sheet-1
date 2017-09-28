

package main

import ( 
	"fmt"
	"unicode/utf8"
)


func main() {
	word := "hello how are you today."
	fmt.Printf("original string : %v \n",word)
	fmt.Printf("reversed string : %v \n",reverseString(word))
}

func reverseString(word string) string {
	wordLen := utf8.RuneCountInString(word)
	runes := make([]rune,wordLen)
	for _,rune:= range word {
	wordLen--
	runes[wordLen]=rune
	}
	return string(runes[wordLen:])
	}