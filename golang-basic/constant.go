package main

import "fmt"

func main() {
	// const firstName string = "Fandi"
	// const lastName = "Hasnur"

	const (
		firstName string = "Fandi"
		lastName         = "Hasnur"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}
