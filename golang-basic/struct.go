package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var fandi Customer
	fmt.Println(fandi)

	fandi.Name = "Fandi Hasnur"
	fandi.Address = "Indonesia"
	fandi.Age = 24
	fmt.Println(fandi)
	fmt.Println(fandi.Name)
	fmt.Println(fandi.Address)
	fmt.Println(fandi.Age)

	alam := Customer{
		Name:    "Alam",
		Address: "Indonesia",
		Age:     20,
	}
	fmt.Println(alam)

	agung := Customer{"Agung", "Indonesia", 22}
	fmt.Println(agung)

	agung.sayHello("Eko")
	alam.sayHello("Eko")
	fandi.sayHello("Eko")
}
