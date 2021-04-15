package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Rofiq)",
			request: "Rofiq",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Rofiq", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rofiq")
		}
	})
	b.Run("Rizky", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rizky")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rofiq")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Rofiq)",
			request:  "Rofiq",
			expected: "Hello Rofiq",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTes(t *testing.T) {
	t.Run("Rofiq", func(t *testing.T) {
		result := HelloWorld("Rofiq")
		require.Equal(t, "Hello Rofiq", result)
	})
	t.Run("Rizky", func(t *testing.T) {
		result := HelloWorld("Rizky")
		require.Equal(t, "Hello Rizky", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before UNIT TEST")

	m.Run()

	fmt.Println("After UNIT TESTING")
}

func SkipTest(t *testing.T) {
	if runtime.GOOS == "Ubuntu" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Rofiq")
	require.Equal(t, "Hello Rofiq", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rofiq")
	require.Equal(t, "Hello Rofiq", result)
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rofiq")
	assert.Equal(t, "Hello Rofiq", result)
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		t.Error("Result must be Hello World")
	}
}

func TestHelloWorldRofiq(t *testing.T) {
	result := HelloWorld("Rofiq")
	if result != "Hello Rofiq" {
		t.Fatal("Result must be 'Hello Rofiq'")
	}
}
