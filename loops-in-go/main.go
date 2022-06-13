package main

import "fmt"

// the code below is used to show the use of loops in go
func main() {
	// the go language only has the for loop

	// the code below is used to print numbers from 1 to 5 using the for loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// the code below is used to create a slice of names of the user
	var names = []string{}

	// the code below is to use the append method for adding the data to the
	// names slice

	// here we are using the append method to add the data to the names slice and
	// in the current scenario we are adding multiple values to the names slice
	names = append(names, "nana", "trisha", "chad", "angela", "helen", "mala", "hitesh")

	// the line of code below is for debugging purpose
	fmt.Println("using normal for each loop")
	// the code below is to use the for each loop for looping over the names slice

	// the loop that we are using below ca also be used with array

	// the for each loop below takes the index and the variable i.e. userName (for
	// refering to the values in the names slice at the particular index ) as input

	// the range keyword is used for creating the for each loop and the names is the
	// name of our slice on which we are iterating
	for index, userName := range names {
		// the code below is used to print the value of userName along with the
		// index to which it belongs
		fmt.Printf("The value at %v index is %v\n", index, userName)
	}

	// the below line of code is for debugging purpose
	fmt.Println("using for each loop without the index variable ")
	// the code below is used to show the use of _ i.e. the blank identifier that
	// we can use in place of the index variable if we do not want to use the index
	// variable in our loop below

	// in the code below we have replaced the index variable with _ underscore i.e.
	// the blank identifier since we do not want to use that index variable
	for _, userName := range names {
		// the code below is used to print the value of userName along with the
		// index to which it belongs
		fmt.Println(userName)
	}

	// the line of code below is for debugging purpose
	fmt.Println("using normal for loop")
	fmt.Println("The values in the slice names are: ")
	// the code below is to use the normal for loop for looping over the names slice
	for index := 0; index < len(names); index++ {
		// the code below is used to print the value in the names slice
		fmt.Println(names[index])
	}

}
