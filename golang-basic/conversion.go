package main

import "fmt"

func main() {
	var nilai32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Fandi Hasnur"
	var f uint8 = name[0]
	var fString = string(f)

	fmt.Println(name)
	fmt.Println(f)
	fmt.Println(fString)
}
