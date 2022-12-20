package utils;
import "fmt";

type Error struct {
	name string
	description string
	line int
}

// NewError: Creates new error object and returns its address (to be assigned to variable)
func NewError(ename, edesc string, eline int) *Error {
	if eline < 0 {
		fmt.Printf("<Line of Error (%d) cannot be Negative>\n\n", eline);
	} else if len(ename) == 0 || len(edesc) == 0 {
		fmt.Printf("<Name of Error/Desc. cannot be Empty>\n\n");
	}

	return &Error{
		name: ename,
		description: edesc,
		line: eline,
	}
}

// PrintError: prints formatted error [.env:<line> <error>: <error-msg>]
func PrintError(e Error) {
	fmt.Printf(".env:%d\n\t%s:\t%s\n", e.line, e.name, e.description);
}
