//PROBLEM THREE : Fizz Buzz print numbers from 1 to 100
//Replacing multiples of 3 with fizz.
//Replacing multiples of 5 with buzz.
//Replacing multiples of 3 and 5 with fizzbuzz.

package main

import (
	"fmt"
)

func main() {
	//for loop that will print the numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		//if i%3 equals 0 it means the number is a multiple of 3
		//if i%5 equals 0 it means the number is a multiple of 5
		//when the mod of a number = 0 their is no remainder
		//this proves that it is a multiple of the number.

		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizz buzz.\n") // prints fizz buzz when the number is both a multiple of 3 and 5.
		} else if i%5 == 0 {
			fmt.Printf("buzz.\n") // prints buzz when the number is only a multiple of 5.
		} else if i%3 == 0 {
			fmt.Printf("Fizz.\n") //prints fizz when the number is only a multiple of 3.
		} else {
			fmt.Printf("%v.\n", i./) //prints the number if it is not a multiple of 3 or 5.
		}
	}
}
