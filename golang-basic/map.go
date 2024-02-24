package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Fandi"
	// person["address"] = "Baubau"

	person := map[string]string{
		"name":    "Fandi",
		"address": "Baubau",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Fandi"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
