package utils

import (
	"strings"
	"regexp"
)

func ValidatePairline(pairline string, line int) {
	key, value, _ := strings.Cut(pairline, "=");

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

	if isKeyValueless(value) {
		var err = Error{
			name: "ValuelessKey",
			description: "<key (" + key + ") is invalid or valueless (" + value + ")>",
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