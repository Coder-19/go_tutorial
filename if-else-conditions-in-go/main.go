package main

import "fmt"

// the code below is to show the use of if else statement in go
func main() {
	// the code below is used to check that if the value of num is divisible by
	// 2 then the number is even else the number is odd

	// the code below is used to create a variable having a value
	var num uint = 14

	// the code below is used to check that if the number is divisible by 2 then
	// the number is even else the number is odd
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}

}
