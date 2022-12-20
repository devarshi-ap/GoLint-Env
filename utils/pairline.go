package utils

import (
	"strings"
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

	
}

func hasDuplicate(key string) (string, bool) {
	if Pairs_Map.containsKey(key) {
		return Pairs_Map[key], true
	} else {
		return "", false;
	}
}