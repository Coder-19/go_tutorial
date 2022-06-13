package main

import "fmt"

// the code below is used to show the use of the break and continue statement
func main() {

	// the code below is for debugging purpose
	fmt.Println("showing use of the break keyword")

	// the code below is used to print the numbers from 1 to 5 but if the value of
	// i is 3 then breaking out of the loop using the break statement
	for i := 1; i <= 5; i++ {

		// going out or breaking out of the loop when i is equal to 3
		if i == 3 {
			break
		}

		// the code below is used to print the value of i to the console
		fmt.Println(i)
	}

	// the code below is for debugging purpose
	fmt.Println("showing use of the continue keyword")

	// the code below is used to print the numbers from 1 to 5 using the for loop
	for i := 1; i <= 5; i++ {

		// skips the iteration when i is equal to 3
		if i == 3 {
			continue
		}

		// the code below is used to print the value of i to the console
		fmt.Println(i)
	}
}
