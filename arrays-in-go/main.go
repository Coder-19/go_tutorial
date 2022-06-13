package main

import "fmt"

// the code below is used to show the use of arrays in golang
func main() {
	// the code below is used to create an array for getting the list of users who
	// have booked the tickets for the conference
	// var bookings = [50]string{"Nana","Jack"}     // this array can store 50 values since the size
	// of the bookings array is 50 and the values that it can store is of type string

	// the code below is used to create a dynamic array for getting the names of the users
	// who have booked the tickets

	// the below array is an empty array of dynamic length
	// var bookings = [...]string{} // here we are using [...] to create an array of
	// string type having dynamic length

	// the code below is used to create an array of size 50 and of type string
	var bookings [50]string

	// the code below is used to create an empty array using the another method with
	// length of array being dynamic
	// var bookings []string

	// the code below is used to set the value at index 1 as nana
	bookings[0] = "Nana"
	// the code below is used to set the value at index 2 as Trisha
	bookings[1] = "Trisha"

	// the code below is to use the fmt.Println() for priting the elements of the
	// bookings array
	fmt.Println("the data in the array is: ")
	fmt.Println(bookings)

	// the code below is used to print the value at index 1 of the bookings array
	fmt.Printf("The value at index 1 of the bookings array is: %v \n", bookings[1])

	// the code below is to use the len() function to print the length of the array
	// bookings
	fmt.Println("The length of the bookings array is:", len(bookings))

}
