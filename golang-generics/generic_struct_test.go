package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Fandi",
		Second: "Hasnur",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Fandi",
		Second: "Hasnur",
	}

	assert.Equal(t, "Alam", data.ChangeFirst("Alam"))
	assert.Equal(t, "Hello Fandi", data.SayHello("Fandi"))
}
