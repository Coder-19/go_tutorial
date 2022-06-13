package main

import "fmt"

// the code below is used to create an anonymous function that retuns an
// anonymous function
func displayNumber() func() int {

	number := 10

	// the code below is used to return an anonymous function which has a return
	// type of int
	return func() int {
		number++
		return number
	}
}

// the code below is to show the use of anonyous functions in go
func main() {
	// the code below is used to create an anonymous function to greet the users
	var greet = func() {
		fmt.Println("Hello from anonymous greet() function")
	}

	// the code below is used to call the greet() function below
	greet()

	// the code below is used to create an anonymous function that takes two numbers
	// and print their sum
	var sum = func(num1 int, num2 int) {
		var add = num1 + num2
		fmt.Println("Sum of number", num1, "and", num2, "is", add)
	}

	// the code below is used to call the sum method below
	sum(2, 4)

	// the code below is used to create a method that takes two numbers from the user
	// and return its product as output
	var multiply = func(num1 float64, num2 float64) float64 {
		return num1 * num2
	}

	// the code below is used to call the multiply method and pass the two numbers and
	// save its product in variable called product
	var product float64 = multiply(2, 6)

	// the code below is used to print the value of product to the console
	fmt.Println("Product", product)

	// the code below is used to call the displayNumber method and save its result in
	// a variable called result
	result := displayNumber()

	// the code below is used to print the value of result anonymous function to the
	// console
	fmt.Println(result())

}
