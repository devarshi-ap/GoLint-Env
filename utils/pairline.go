package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidatePairline(pairline string, line int) {
	key, value, _ := strings.Cut(pairline, "=");
	fmt.Println("key", key, "\nvalue", value);

	if curr_val, b := hasDuplicate(key); b {
		var err = Error{
			name: "DuplicateKey",
			description: "<duplicate key (" + key + ") has existing value (" + curr_val + ")>",
			line: line,
		}
		Error_Stack.Push(err);
	} else {
		// add key-value pair to map
		Pairs_Map[key] = value;
	}

	if len(strings.TrimSpace((key))) == 0 {
		var err = Error{
			name: "EmptyKey",
			description: "<key is empty>",
			line: line,
		}
		Error_Stack.Push(err);
	}

	if isKeyValueless(value) {
		var err = Error{
			name: "ValuelessKey",
			description: "<key (" + key + ") is invalid or valueless (" + value + ")>",
			line: line,
		}
		Error_Stack.Push(err);
	}

	if !hasCorrectDelimiter(key) {
		var err = Error{
			name: "IncorrectDelimiter",
			description: "<key (" + key + ") has invalid delimiter; key must be uppercase & underscore-delimited>",
			line: line,
		}
		Error_Stack.Push(err);
	}

	if !hasCorrectLeadingChar(key) {
		var err = Error{
			name: "IncorrectLeadingChar",
			description: "<key (" + key + ") has invalid leading character; must start with underscore or uppercase letter>",
			line: line,
		}
		Error_Stack.Push(err);
	}

	if hasLowercaseKey(key) {
		var err = Error{
			name: "LowercaseInKey",
			description: "<key (" + key + ") has lowercase character; must be uppercase>",
			line: line,
		}
		Error_Stack.Push(err);
	}
}

func hasDuplicate(key string) (string, bool) {
	if Pairs_Map.containsKey(key) {
		return Pairs_Map[key], true
	} else {
		return "", false;
	}
}

func isKeyValueless(val string) bool {
	if len(val) >= 3 && ((string(val[0]) == "'" && string(val[len(val)-1]) == "'") || (string(val[0]) == "\"" && string(val[len(val)-1]) == "\"")) {
		return false;
	}
	var rgx = regexp.MustCompile(`^[^a-zA-Z0-9]+$`)
	new_str := rgx.ReplaceAllString(val, ``)
	return len(strings.TrimSpace(new_str)) == 0;
}

func hasCorrectDelimiter(key string) bool {
	match, _ := regexp.MatchString(`^[A-Z_]+$`, key);
	return match;
}

func hasCorrectLeadingChar(key string) bool {
	firstChar := string(key[0]);
	match, _ := regexp.MatchString(`^[A-Za-z_]+$`, firstChar);
	return match;
}

func hasLowercaseKey(key string) bool {
	match, _ := regexp.MatchString(`[a-z]{1,}`, key);
	return match;
}