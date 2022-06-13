package main

import "fmt"

// the code below is used to show how to use map data structure in go
func main() {
	// the code below is used to create a map a having key type as string and
	// value type as string
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}

	// the code below is used to create a map b having key type as string and value
	// type as int
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
