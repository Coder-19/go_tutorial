package main

import "fmt"

// the for loop can be used as a while loop or the if block

// the code below is used to show how to use the for loop as the while loop
func main() {
	// the below line of code is for debugging purpose
	fmt.Println("Using for loop as while loop")

	// the code below is used to create a variable for storing the value of num
	num := 0

	// the code below is used to run the for loop till the value of num is 10
	for num <= 10 {
		fmt.Println(num)
		num++
	}
}
