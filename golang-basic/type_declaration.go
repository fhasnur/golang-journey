package main

import "fmt"

func main() {
	type NoKTP string

	var ktpFandi NoKTP = "123456789"

	var contoh string = "222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpFandi)
	fmt.Println(contohKtp)
}
