package main

import "fmt"

// the code below is used to show the different ways of declaring the variables
// in go
func main() {
	// the code below is used to create a variable of type uint (i.e. unsigned integer) in go haivng value 50
	var num uint = 50

	// the code below is used to create a variable of type string for getting the
	// first name of the user
	var firstName string = "Programmer"

	// the code below is used to create a variable using := for getting the last name of the user
	lastName := "Coder"

	// the code below is used to create an unitialized variable for getting the
	// age of the user
	var age int

	// assigning the value to the age variable
	age = 12

	// the code below is to use the fmt.Println() method to print a message to the console
	// using string interpolation
	fmt.Println("Hello", firstName, lastName, "how are you?")
	// the code below is to use the fmt.Printf() method to print a formatted output
	// displaying the age of the user
	fmt.Printf("Your age is %v and the number selected by you is %v", age, num) // here first %v stands for the refernce of the age variable and the
	// second %v stands for the num variable
}
