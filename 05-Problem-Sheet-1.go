package main

import (
	"fmt"
	"math/rand"
	"time"
)

//PROBLEM FIVE : Guessing Game
func main() {
	var complete bool = false                        //test clause to exit loop.
	var tries, preinput, input int = 0, 0, 0         //number of tries, users previous input to test for repeats, user input.
	var seed = rand.NewSource(time.Now().UnixNano()) //creats a seed for the rand fucntion based on current time.
	var myRand = rand.New(seed)                      //allows us to creat a number based of the seed.
	var target int = myRand.Intn(1000)               //generates a random number between 0 - 1000.

	//main loop repeats until the user guesses the correct anwser.
	for complete != true {

		fmt.Printf("Enter your guess.\n")
		fmt.Scan(&input)

		//if the user guesses the correct number we output number of tries and exit loop.
		if input == target {
			tries++
			fmt.Printf("You guessed correctly it took you %v tries.\n", tries)
			complete = true
			//if the user guess to high they are told and number of tries is incremented.
		} else if input > target {
			fmt.Printf("your guess was to high\n")
			if preinput != input {
				tries++
				preinput = input
			}
			//if the user guess to low they are told and number of tries is incremented.
		} else if input < target {
			fmt.Printf("your guess was to low\n")
			if preinput != input {
				tries++
				preinput = input
			}
		}

	}
}
