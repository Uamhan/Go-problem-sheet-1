package main

import ( 
	"fmt"
)

//PROBELEM 7 : test a word to see if its a palindrome
func main() {

	//takes user input and passes to the isPal Function
	var testWord string

	fmt.Println("Enter your word to be tested as a palindrome")
	fmt.Scan(&testWord)
	isPal(testWord)
}

//isPal function
//funtion takes in string named word
//gets length of word and stores as wordLen
//loop then checks weather the first and last letters of word are !=
//if so loop marks the passed variable as false
//prints weather the word is a palidrome or not
func isPal(word string) {
	var wordLen int
	wordLen = len(word)
	passed := true
	var i int = 0
	for wordLen > 1 {
		if word[0+i] != word[(len(word)-1)-i] {
			passed = false
		}
		wordLen -= 2
		i++
	}
	if passed == true {
		fmt.Println("your word is a palindrome")
	} else {
		fmt.Println("your word is not a palindrome")
	}
}