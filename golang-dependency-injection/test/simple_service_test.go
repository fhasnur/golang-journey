package test

import (
	"fmt"
	"golang-restful-api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
