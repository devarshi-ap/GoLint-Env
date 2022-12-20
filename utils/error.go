package utils;
import "fmt";

type Error struct {
	name string
	description string
	line int
}

// NewError: Creates new error object and returns its address (to be assigned to variable)
func NewError(ename, edesc string, eline int) *Error {
    return &Error{
		name: ename,
		description: edesc,
		line: eline,
	}
}

// PrintError: prints formatted error [.env:<line> <error>: <error-msg>]
func PrintError(e Error) {
	fmt.Printf(".env:%d %s: %s\n", e.line, e.name, e.description);
}
