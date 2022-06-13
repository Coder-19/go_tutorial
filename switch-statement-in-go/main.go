package main

import "fmt"

// the code below is used to show the use of switch statement in go

// there is no break  statement after the case block in the switch statement

func main() {

	// the property below is used to get the day number from the user
	dayOfWeek := 3

	// the switch statement below takes the dayOfWeek as input condition and then
	// match that according to the cases present
	switch dayOfWeek {

	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	// the code below is used to create the default case
	default:
		fmt.Println("Invalid day")
	}
}
