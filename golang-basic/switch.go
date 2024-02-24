package main

import "fmt"

func main() {
	name := "Fandi"

	switch name {
	case "Fandi":
		fmt.Println("Hello Fandi")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	name = "Hasnur"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
