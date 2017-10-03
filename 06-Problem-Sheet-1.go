package main

import (
	"fmt"
)

//PROBLEM 6 : pass a list to a function that returns the minimum and maximum value from the list.
func main() {

	//Creates a slice with value 5,1,4,12,2,7.
	x := []int{5, 1, 4, 12, 2, 7}

	//gets the min and max values of the slice from bigSmall function commented on below.
	min, max := bigSmall(x)
	//prints min and max values.
	fmt.Printf("min : %v, max : %v", min, max)
}

//BigSmall
//takes in a slice and gets the length of this slice and sets the min and max values to the first item in the list.
//uses a forloop to iterate through the list and tests each item against the min and max values.
//if item is less than the min value it will replace it.
//if item is greater than max value it will replace it.
//when loop is complete min max values will be correct.
//returns min and max values.
func bigSmall(myList []int) (int, int) {

	var min, max int = myList[0], myList[0]
	var x = len(myList)
	for i := 0; i < x; i++ {

		if myList[i] < min {
			min = myList[i]
		}
		if myList[i] > max {
			max = myList[i]
		}
	}
	return min, max

}
