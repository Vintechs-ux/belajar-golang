package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type iniTest struct {
	name   string
	req    string
	exp    string
	result string
}

func BenchmarkHelloWorldTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Siregar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Siregar")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name string
		req  string
	}{
		{
			name: "HelloWorld(eko)",
			req:  "Eko",
		},
		{
			name: "HelloWorld(siregar)",
			req:  "Siregar",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.req)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Menghubungkan ke PostgreSQL di port 5000")

	m.Run()

	// after
	fmt.Println("Memutus koneksi database... ")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello World Eko", result, "Result must be Hello World Eko")
}

func TestHelloWorldSiregar(t *testing.T) {
	result := HelloWorld("Siregar")
	assert.Equal(t, "Hello World Siregar", result, "Result must be 'Hello World Siregar'")
}

func TestHelloWorldWindows(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Unit test tidak bisa dijalankan oleh linux")
	}
	result := HelloWorld("Sitompul")
	assert.Equal(t, "Hello World Sitompul", result, "Result must be Hello World Sitompul")
}

func TestHelloWorldSub(t *testing.T) {
	t.Run("Jecky", func(t *testing.T) {
		result := HelloWorldSub("Jecky")
		require.Equal(t, "Hello World Sub Jecky", result, "Result must be 'Hello World Sub Jecky'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		req      string
		expected string
		result   string
	}{
		{
			name:     "Test dengan nama Joko",
			req:      "Joko",
			expected: "Hello World Joko",
			result:   "Result must be 'Hello World Joko'",
		},

		{
			name:     "Test dengan nama Parsya",
			req:      "Parsya",
			expected: "Hello World Parsya",
			result:   "Result must be 'Hello World Parsya'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.req)
			assert.Equal(t, test.expected, result, test.result)
		})
	}
}
