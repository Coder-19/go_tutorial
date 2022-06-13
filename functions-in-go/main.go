package main

import "fmt"

// the code below is used to create a method for greeting the user
func greet() {
	fmt.Println("Good Morning")
}

// the code below is used to create a method that takes two numbers as input and
// add them up

// the addNumbers() method below takes n1 and n2 of type int as input and then add them
func addNumbers(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum:", sum)
}

// the code below is used to create a method that will multiply two numbers and
// return their product as output

// the method below takes two numbers n1 and n2 of type int as input and the return
// type of the method is int
func multiplyNumbers(n1 int, n2 int) int {

	// the code below is used  to multiplty n1 with n2 and save their result in the
	// mul variable
	mul := n1 * n2

	// the code below is used to return the value of mul as output
	return mul
}

// the code below is used to create a method that takes n1 and n2 numbers of type
// int as input and then do calculation on them and then return the product
// and difference of the n1 and n2

// the return type of the method below is of type (int,int)
func calculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2

	// return two values

	// the return statement below first returns the sum and then retuns the difference
	return sum, difference
}

// the code below is used to show how functions are created and used in go
func main() {
	// the code below is used to call the greet() message
	greet()

	// the code below  is used to create property for getting the value of num1
	// and num2
	var num1 int = 2
	num2 := 4

	// the code below is used to call the addNumbers() method
	addNumbers(num1, num2)

	// the code below is used to call the mulNumbers() method and pass num1 and
	// num2 as input and then save the value returned by this method in a variable \
	// called product
	var product int = multiplyNumbers(num1, num2)

	// the code below is used to print the value of product to the console
	fmt.Println("Product:", product)

	// the code below is used to create two variables for getting the sum and
	// difference value as output from the calculate method

	// the order of creating the sum and difference variable below is important since
	// the calculate method first retuns the sum and then the difference of the numbers
	// num 1 and num 2
	sum, difference := calculate(num1, num2)

	// the code below is used to print the sum and difference of num1 and num2 to the
	// console
	fmt.Printf("The sum of %v and %v is %v and the difference between %v and %v is %v \n", num1, num2, sum, num1, num2, difference)
}
