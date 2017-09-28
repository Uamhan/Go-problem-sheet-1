package main

import (
	"fmt"
	"math"
)

func main() {

	x := 9.0                                           // the value we want to get the square root off
	fmt.Printf("math.Sqrt of x = %v \n", math.Sqrt(x)) //prints the math.sqrt result of x
	//for loop that prints the retun value of the z_next function untill z = z_next
	for z := 1.0; z != z_next(x, z); z = z_next(x, z) {
		fmt.Printf("x = %v, z = %v\n", x, z)
	}
}

//function that takes in the number you want the square root of and our z value and returns
//the value givin by newtons method for square roots.
func z_next(x, z float64) float64 {
	return (z - ((z*z - x) / (2 * z)))
}
