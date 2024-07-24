package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Fandi",
		MiddleName: "Meylwan",
		LastName:   "Hasnur",
		Age:        25,
		Married:    false,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
