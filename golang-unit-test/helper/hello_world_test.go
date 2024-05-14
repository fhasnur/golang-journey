package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Fandi")
	require.Equal(t, "Hello Fandi", result, "Result must be 'Hello Fandi'")
}

func TestHelloWordlRequire(t *testing.T) {
	result := HelloWorld("Fandi")
	require.Equal(t, "Hello Fandi", result, "Result must be 'Hello Fandi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWordlAssert(t *testing.T) {
	result := HelloWorld("Fandi")
	assert.Equal(t, "Hello Fandi", result, "Result must be 'Hello Fandi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

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
