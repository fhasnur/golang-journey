package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldFandi(t *testing.T) {
	result := HelloWorld("Fandi")

	if result != "Hello Fandi" {
		// error
		t.Error("Result must be 'Hello Fandi'")
	}

	fmt.Println("TestHelloWorldFandi Done")
}

func TestHelloWorldHasnur(t *testing.T) {
	result := HelloWorld("Hasnur")

	if result != "Hello Hasnur" {
		// error
		t.Fatal("Result must be 'Hello Hasnur'")
	}

	fmt.Println("TestHelloWorldHasnur Done")
}
