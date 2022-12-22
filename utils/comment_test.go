package utils

import (
	"testing"
)

func TestIsCommentEmpty(t *testing.T) {
	cases := []struct {
        input          string
        expectedOutput bool
    }{
        {"#", true},
        {"# ", true},
        {"#     ", true},
        {"#to pimp a butterfly", false},
        {"# to# pimp# a# butterfly#", false},
        {"#  #", false},
    }

    for _, c := range cases {
        if output := isCommentEmpty(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.input, c.expectedOutput, output)
        }
    }
}

func TestCommentHasTrailingSpace(t *testing.T) {
	cases := []struct {
        input          string
        expectedOutput bool
    }{
        {"#", false},
        {" #to pimp a butterfly", false},
        {"#  #", false},
        {"# to# pimp# a# butterfly# ", true},
        {"# ", true},
        {"#     ", true},
    }

	for _, c := range cases {
		if output := commentHasTrailingSpace(c.input); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.input, c.expectedOutput, output)
		}
	}
}

func TestIsCommentIncorrectlySpaced(t *testing.T) {
	cases := []struct {
        input          string
        expectedOutput bool
    }{
        {"#", true},
        {"# to pimp a butterfly", true},
        {"# to# pimp# a# butterfly# ", true},
        {"# #", true},
        {"#  x", false},
        {"#  #   ", false},
    }

	for _, c := range cases {
		if output := isCommentCorrectlySpaced(c.input); output != c.expectedOutput {
			t.Errorf("incorrect output for `%s`: expected `%t` but got `%t`", c.input, c.expectedOutput, output)
		}
	}
}