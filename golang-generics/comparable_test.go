package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Fandi", "Fandi"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](true, true))
}
