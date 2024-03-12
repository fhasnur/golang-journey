package main

import "fmt"

func main() {
	firstName := "Fandi"
	lastName := "Hasnur"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
