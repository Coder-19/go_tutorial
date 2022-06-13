package main

import "fmt"

// the code below is used to get the input from the user and display it in the console
func main() {
	// the property below is used to get the name of the user
	var firstName string

	// the property below is used to get the lastName of the user
	var lastName string

	// the property below is used to get the email of the user
	var email string

	// the property below is used to get the number of tickets that the user wants to buy
	var userTickets int

	// the code below is used to initialize the name of the user
	// userName = "Tom"

	// the code below is used to initialize the userTickets property so as to know how many tickets the
	// user wants to buy
	// userTickets = 10

	// the code below is used to print the address of the username variable in the memory
	// fmt.Println("the address where userName variable is present in memory is: ", &userName)

	// the code below is used to print the address of the userTickets variable in the memory
	// fmt.Println("the address where the userTickets variable is stored in the memory is: ", &userTickets)

	// the code below is used to display a message to the user for entering the username
	fmt.Println("Please enter your first name")
	// the code below is to use the fmt.Scan() function for getting the data from the user and the fmt.Scan()
	// function takes the variable name as input in which we want to store the data that we get from the user

	// in the code below we are using the &userName for getting the address of the variable userName for storing the
	// data in the userName variable in the memory
	fmt.Scan(&firstName) // here we are getting the name of the user as input and saving that in a variable
	// named userName

	// the code below is used to print a message for the user to enter the lastName
	fmt.Println("Please enter your last name")

	// the code below is to use the fmt.Scan() method for getting the lastName of the user
	// and then storing that value in the lastName variable
	fmt.Scan(&lastName)

	// the code below is used to create a message for the user to enter the email
	fmt.Println("Please enter your email address")

	// the code below is to use the fmt.Scan() method for getting the email of the user
	// and storing it in email variable
	fmt.Scan(&email)

	// the code below is used to display a message to the user for entering the number of tickets that the user
	// wants
	fmt.Println("Please enter the number of tickets")

	// in the code below we are using the &userTickets for getting the address of the variable i.e. where the
	// variable userTickets is present in the memory for storing the value of the tickets entered by the user
	// in the memory
	fmt.Scan(&userTickets) // here we are getting the userTickets as input i.e. the number of tickets that
	// the user wants and then storing that in a variable named userTickets

	// the code below is used to print a confirmation message to the user
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
