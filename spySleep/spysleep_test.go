package main

import (
	"bytes"
	"reflect"
	"testing"
)

type testSleep struct {
	Calls []string
}

func (t *testSleep) Sleep() {
	t.Calls = append(t.Calls, "sleep")
}

func (t *testSleep) Write(p []byte) (n int, err error) {
	t.Calls = append(t.Calls, "write")
	return len(t.Calls), nil
}

func TestPrintSleep(t *testing.T) {

	t.Run(`Prints 1, 2, 3, Go!`, func(t *testing.T) {
		b := &bytes.Buffer{}
		ts := &testSleep{}

		printAndSleep(b, ts)

		got := b.String()
		expected := `3
2
1
Go!`
		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("Prints in correct order", func(t *testing.T) {
		ts := &testSleep{}

		printAndSleep(ts, ts)

		expected_order := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}

		if !reflect.DeepEqual(expected_order, ts.Calls) {
			t.Errorf("It should call sleep 3 times")
		}
	})

}
