package main

import "fmt"

// the code below is used to show the use of structures in go
func main() {
	// the code below is used to create a user structure
	type User struct {
		// creating a name property of type string
		name string
		// creating an occupation property of type string
		occupation string
		// creating age property of type int
		age int
	}

	// the code below is used to create a new user from the user structure
	u := User{
		name:       "Angela",
		occupation: "Software Developer",
		age:        30,
	}

	// the below line of code is for debugging purpose to print the value of u
	// to the console
	fmt.Println(u)

	// the code below is to print the name of the user u to the console
	fmt.Println(u.name)

	// the code below is used to create anonymous structures in go
	person := struct {
		name        string
		age         int
		designation string
	}{
		name:        "Chad",
		age:         40,
		designation: "Team Lead",
	}

	// the code below is used to print the value of person to the console
	fmt.Println(person)

	// the code below is used to print the designation of person to the console
	fmt.Println(person.designation)

	// the code below is used to create the nested struct in go

	// the code below is used to create the address struct
	type Address struct {
		city    string
		country string
	}

	// the code below is used to create the developer struct
	type Developer struct {
		name    string
		age     int
		Address // using the address struct as a nested struct to developer struct
	}

	p := Developer{
		name: "John Doe",
		age:  34,
		Address: Address{
			city:    "New York",
			country: "USA",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)
	fmt.Println("Country:", p.country)

	// Creating structure
	type Student struct {
		name   string
		branch string
		year   int
	}

	// Creating nested structure
	type Teacher struct {
		name    string
		subject string
		exp     int
		details Student
	}

	// Initializing the fields
	// of the structure
	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{"Bongo", "CSE", 2},
	}

	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", result.details.name)
	fmt.Println("Student's branch name: ", result.details.branch)
	fmt.Println("Year: ", result.details.year)

	// the code below is used to create the slice of the User structure
	users := []User{}

	// the code below is to use the append() method to add data to the user structure
	users = append(users, User{"John Doe", "gardener", 20})
	users = append(users, User{"Roger Roe", "driver", 25})
	users = append(users, User{"Paul Smith", "programmer", 23})
	users = append(users, User{"Lucia Mala", "teacher", 40})
	users = append(users, User{"Patrick Connor", "shopkeeper", 40})
	users = append(users, User{"Tim Welson", "programmer", 28})
	users = append(users, User{"Tomas Smutny", "programmer", 30})

	// using the for each loop for printing the data inside the users slice to the
	// console

	fmt.Println("The data inside the user slice is:")
	for _, user := range users {
		fmt.Println(user)
	}

	// the code below is used to create a persons slice of type Student and then adding data to
	// that slice
	var persons = []Student{
		{
			name:   "Chad",
			branch: "Computer Science and Engineering",
			year:   2,
		},
		{
			name:   "Mala",
			branch: "Computer Science and Engineering",
			year:   2,
		},
		{
			name:   "Helen",
			branch: "Computer Science and Engineering",
			year:   2,
		},
	}

	// the code below is to use the for loop for printing the data inside the
	// person slice to the console
	fmt.Println("The names of the students are:")
	for index := 0; index < len(persons); index++ {
		fmt.Println(persons[index].name)
	}

}
