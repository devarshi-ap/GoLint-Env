package utils

import (
	"fmt"
)

type Stack []Error

var Error_Stack Stack

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Print: print stack elements
func (s *Stack) Print() {

	if len(*s) == 0 {
		// No Error YAYY!
		fmt.Println(colorPurple, "\n", "✨ No Errors Found ✨");
		DrawMoose();
	} else {
		fmt.Println(colorYellow, "\n", "🚧", len(Error_Stack), "Errors Caught 🚧")
		fmt.Println("-----------------------");
		for i := 0; i < len(*s); i++ {
			PrintError((*s)[i])
		}
		fmt.Println();
	}
}

// Push: append new value @ end of stack (check if prev key is lexicographically smaller- in order)
func (s *Stack) Push(e Error) {
	*s = append(*s, e)
}