package main

import "fmt"

// Interface kosong: interface{} / any
func Ups() any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
