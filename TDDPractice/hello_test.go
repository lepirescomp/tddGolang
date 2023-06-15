package TDDPractice

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Test Hello Default Case", func(t *testing.T) {
		got := Hello("", "Test")
		expected := "Hello, Test"
		assertCorrectMessage(t, got, expected)

	})
	t.Run("Test Spanish Case", func(t *testing.T) {
		got := Hello("spanish", "Test")
		expected := "Hola, Test"
		assertCorrectMessage(t, got, expected)
	})

}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func ExampleHello() {
	result := Hello("", "TestName")
	fmt.Print(result)
	// Output: Hello, TestName
}

func BenchmarkNestedLoop(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		NestedLoop()
	}
}
