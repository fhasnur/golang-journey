package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fandi")

	if result != "Hello Fandi" {
		// error
		panic("Result is not 'Hello Fandi'")
	}
}
