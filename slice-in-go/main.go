package main

import "fmt"

// the code below is used to show the use of slice in go
func main() {
	// slice is basically an abstraction of arrays in go since we cannot define
	// arrays with dynamic size in go

	// the code below is used to create an empty slice
	// var bookedList  = []string{}

	// the code below is used to create a dynamic slice
	// bookedList := []string{}

	// the code below is used to create a slice of type string
	var bookedList []string

	// the code below is to use the append method to add a value to the bookedList
	// slice

	// the code below is used to add nana to the slice bookedList

	// in the code below we are adding only 1 value to the slice
	bookedList = append(bookedList, "Nana")

	// the code below is used to add Trisha, Helen to the bookedList slice

	// in the code below we are adding two values to the slice
	bookedList = append(bookedList, "trisha", "Helen")

	// the code below is used to print the slice in the console
	fmt.Println("The data inside the bookList slice is:", bookedList)

	fmt.Println("the data in the slice is: ")
	fmt.Println(bookedList)

	// the code below is used to print the value at index 1 of the bookings array
	fmt.Printf("The value at index 1 of the bookings slice is: %v \n", bookedList[1])

	// the code below is to use the len() function to print the length of the array
	// bookings
	fmt.Println("The length of the bookings slice is:", len(bookedList))

}
