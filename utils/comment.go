package utils

import (
	"strings"
	"unicode"
	"regexp"
)

/*
Incorrectly-Spaced Comments - comments should be preceded with a space
	❌ Wrong
	#comment
	✅ Correct
	# comment

Trailing Whitespace - detects if a line has trailing whitespace
	❌ Wrong
	# comment<whitespace>

	✅ Correct
	# comment

Empty Comment (unnecessary) - detects if line only has '#'
	❌ Wrong
	#
	#<whitespace>

	✅ Correct
	# <non-whitespace>
*/

func Validate(comment string, line int) {
	if isCommentEmpty(comment) {
		var err = Error{
			name: "EmptyComment",
			description: "<comment has no content>",
			line: line,
		}
		Error_Stack.Push(err);
	}

	if commentHasTrailingSpace(comment) {
		var err = Error {
			name: "CommentTrailingSpace",
			description: "<comment has trailing space>",
			line: line,
		}
		Error_Stack.Push(err);
	}

	if isCommentIncorrectlySpaced(comment) {
		var err = Error {
			name: "CommentIncorrectlySpaced",
			description: "<comment must follow [#<space><non-space>] format>",
			line: line,
		}
		Error_Stack.Push(err);
	}
}

func isCommentEmpty(comment string) bool {
	return len(strings.TrimSpace(comment)) == 1
}

func commentHasTrailingSpace(comment string) bool {
	lastChar := comment[len(comment) - 1];
	return unicode.IsSpace(rune(lastChar));
}

func isCommentIncorrectlySpaced(comment string) bool {
	if isCommentEmpty(comment) {
		return false;
	} else {
		match, _ := regexp.MatchString(`(^#[ ][^ ]+$)`, comment);
		return !match;
	}
}