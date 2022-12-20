package utils;

type Stack []Error;

var Error_Stack Stack;

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Print: print stack elements
func (s *Stack) Print() {
	for i:=0; i<len(*s); i++ {
		PrintError((*s)[i]);
	}
}

// Push: append new value @ end of stack
func (s *Stack) Push(e Error) {
	*s = append(*s, e);
}