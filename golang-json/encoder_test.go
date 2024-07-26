package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Fandi",
		MiddleName: "Meylwan",
		LastName:   "Hasnur",
	}

	encoder.Encode(customer)
}
