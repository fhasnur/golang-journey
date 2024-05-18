package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Fandi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fandi")
		}
	})
	b.Run("Hasnur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hasnur")
		}
	})
}

func BenchmarkHelloWorldFandi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fandi")
	}
}

func BenchmarkHelloWorldHasnur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hasnur")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fandi",
			request:  "Fandi",
			expected: "Hello Fandi",
		},
		{
			name:     "Meylwan",
			request:  "Meylwan",
			expected: "Hello Meylwan",
		},
		{
			name:     "Hasnur",
			request:  "Hasnur",
			expected: "Hello Hasnur",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Fandi", func(t *testing.T) {
		result := HelloWorld("Fandi")
		require.Equal(t, "Hello Fandi", result, "Result must be 'Hello Fandi'")
	})
	t.Run("Hasnur", func(t *testing.T) {
		result := HelloWorld("Hasnur")
		require.Equal(t, "Hello Hasnur", result, "Result must be 'Hello Hasnur'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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
