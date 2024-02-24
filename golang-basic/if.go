package main

import "fmt"

func main() {
	name := "Fandi"

	if name == "Fandi" {
		fmt.Println("Hello Fandi")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
