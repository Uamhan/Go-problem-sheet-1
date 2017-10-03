//PROBLEM FOUR : Write a function that calculates the sum of the digits of a factorial.
package main

import (
	"fmt"
)

func main() {

	addFact(10) // calls fucntion addFact.

}

//First loop calculates the factorial of the argument by multiplying the argument against its self then taking away one until its no longer greater than 0
//second loop gets mod 10 of the factorial and adds it to the sum and is then divded by ten this adds each individual number to the total sum
//the sum is then printed
func addFact(x int) {
	var fact, sum int = 1, 0

	for i := x; i > 0; i-- {
		fact *= i
	}
	for fact > 0 {
		sum += fact % 10
		fact /= 10
	}
	fmt.Printf("%v.\n", sum)
}
