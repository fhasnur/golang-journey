package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Fandi",
	}
	second := map[string]string{
		"Name": "Fandi",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Fandi"}
	second := []string{"Fandi"}

	assert.True(t, slices.Equal(first, second))
}
