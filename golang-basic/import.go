package main

import (
	"fmt"
	"golang-basic/helper"
)

func main() {
	result := helper.SayHello("Fandi")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Fandi")) // tidak bisa diakses
}
