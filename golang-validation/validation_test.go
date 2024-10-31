package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "fandi"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}
