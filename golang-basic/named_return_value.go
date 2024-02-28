package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Fandi"
	middleName = "Meylwan"
	lastName = "Hasnur"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
