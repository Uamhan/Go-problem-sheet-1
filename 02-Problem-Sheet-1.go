//PROBLEM TWO : print current time.
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()                        // sets varible t to the current time.
	fmt.Printf("It is currently %v.\n", t) // prints varible t
}
