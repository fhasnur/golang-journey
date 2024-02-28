package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Fandi")
	fmt.Println(result)

	fmt.Println(getHello("Agung"))
	fmt.Println(getHello("Alam"))
}
