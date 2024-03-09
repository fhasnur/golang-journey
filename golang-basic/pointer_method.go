package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fandi := Man{"Fandi"}
	fandi.Married()

	fmt.Println(fandi.Name)
}
