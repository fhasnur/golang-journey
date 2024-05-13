package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldFandi(t *testing.T) {
	result := HelloWorld("Fandi")

	if result != "Hello Fandi" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorldFandi Done")
}

func TestHelloWorldHasnur(t *testing.T) {
	result := HelloWorld("Hasnur")

	if result != "Hello Hasnur" {
		// error
		t.FailNow()
	}

	fmt.Println("TestHelloWorldHasnur Done")
}
