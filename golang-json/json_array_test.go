package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Fandi",
		MiddleName: "Meylwan",
		LastName:   "Hasnur",
		Hobbies:    []string{"Pen Spinning, Coding, Parkour"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fandi","MiddleName":"Meylwan","LastName":"Hasnur","Age":0,"Married":false,"Hobbies":["Pen Spinning, Coding, Parkour"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}
