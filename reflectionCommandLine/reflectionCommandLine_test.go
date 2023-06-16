package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintOut(t *testing.T) {
	b := &bytes.Buffer{}

	printOut(b, "a")

	got := b.String()
	expected := "a"

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestInstructions(t *testing.T) {

	b := &bytes.Buffer{}
	commandLine := CommandLineTool{w: b}

	commandLine.instructions()

	got := b.String()
	var result []string
	for _, item := range got {
		result = append(result, string(item))
	}

	expectedFirstLine := "Tool"
	expectedSecondLine := "Enter"

	assertContains(t, result, expectedFirstLine)
	assertContains(t, result, expectedSecondLine)

}

type userInputMock struct {
	ScanCalls int
	TextCalls int
}

func (u *userInputMock) Scan() bool {
	u.ScanCalls += 1
	if u.ScanCalls == 2 {
		return false
	}
	return true
}

func (u *userInputMock) Text() string {
	u.TextCalls += 1
	return ""
}

func TestInteract(t *testing.T) {

	t.Run("Check messages showed to user", func(t *testing.T) {
		b := &bytes.Buffer{}
		c := CommandLineTool{
			w: b,
		}

		u := userInputMock{}

		interact(&c, &u)

		message := b.String()
		result := strings.Split(message, "\n")

		expectedFirstLine := "Method1"
		expectedSecondLine := "Method2"
		expectedThirdLine := "ShowCommands"

		assertContains(t, result, expectedFirstLine)
		assertContains(t, result, expectedSecondLine)
		assertContains(t, result, expectedThirdLine)
	})

	t.Run("Check scan and text methods", func(t *testing.T) {
		b := &bytes.Buffer{}
		c := CommandLineTool{
			w: b,
		}

		u := userInputMock{}

		interact(&c, &u)

		if u.ScanCalls != 2 {
			t.Errorf("It should call Scan 2 time")
		}

		if u.TextCalls != 1 {
			t.Errorf("It should call Text 1 time")
		}
	})

}

func TestCommandLineTool(t *testing.T) {

	cases := []struct {
		Name           string
		Input          int
		ExpectedOutput string
	}{
		{"method1",
			0,
			"Method1"},
		{"method2",
			1,
			"Method2"},
		{"showComands",
			1,
			"Method1"},
		{
			"Doesn't exists",
			4,
			"exists",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			b := &bytes.Buffer{}
			c := CommandLineTool{
				w: b,
			}
			c.executeMethod(test.Input)

			message := b.String()
			result := strings.Split(message, "\n")

			assertContains(t, result, test.ExpectedOutput)

		})
	}
}

func assertContains(t testing.TB, s []string, need string) bool {
	t.Helper()

	for _, item := range s {
		if need == item {
			return true
		}
	}
	return false
}
