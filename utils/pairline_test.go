package utils
import (
	"testing"
)

func TestIsKeyValueless(t *testing.T) {
	cases := []struct {
		value			string
        expectedOutput 	bool
    }{
        {"", true},
		{"=", true},
		{"==$==(==&&'", true},
		{"'==$==(==&&'", false},
		{"Kendrick Lamar", false},
		{"<123>", false},
		{"'123'", false},
    }

	for _, c := range cases {
		if output := isKeyValueless(c.value); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.value, c.expectedOutput, output)
		}
	}
}

func TestHasCorrectDelimiter(t *testing.T) {
	cases := []struct {
        key          	string
        expectedOutput 	bool
		}{
		{"FOO", true},
		{"FOO_BAR_NOT", true},
		{"foo_BAR", false},
		{"f.b", false},
		{"Foo>bar", false},
		{"FOO BAR", false},
        {"FOO-BAR", false},
    }

	for _, c := range cases {
		if output := hasCorrectDelimiter(c.key); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.key, c.expectedOutput, output)
		}
	}
}

func TestHasCorrectLeadingChar(t *testing.T) {
	cases := []struct {
        key          	string
        expectedOutput 	bool
		}{
		{"  FOO", false},
		{".FOO", false},
		{"*FOO_BAR", false},
		{"1Foo", false},
		{"_FOO", true},
        {"FOO", true},
		{"FOO_BAR", true},
    }

	for _, c := range cases {
		if output := hasCorrectLeadingChar(c.key); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.key, c.expectedOutput, output)
		}
	}
}

func TestHasLowercaseKey(t *testing.T) {
	cases := []struct {
        key          	string
        expectedOutput 	bool
		}{
		{"FOo_BAR", true},
		{".foo", true},
		{"foo_bar", true},
		{"1foo", true},
		{"_FOO", false},
        {"FOO", false},
		{"FOO_BAR", false},
    }

	for _, c := range cases {
		if output := hasLowercaseKey(c.key); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.key, c.expectedOutput, output)
		}
	}
}

func TestHasValidEqualSpacing(t *testing.T) {
	cases := []struct {
        key          	string
        expectedOutput 	bool
		}{
		{"FOO =BAR", false},
		{"FOO_BAR= BAR", false},
		{"FOO = BAR", false},
		{"FOO=BAR", true},
		{"FOO_BAR=BAR", true},
    }

	for _, c := range cases {
		if output := hasValidEqualSpacing(c.key); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.key, c.expectedOutput, output)
		}
	}
}