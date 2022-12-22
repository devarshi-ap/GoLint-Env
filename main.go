package main;
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"golint-env/utils"
);

const filename string = "./assets/.env";

func main() {
	envFile, err := os.Open(filename);
	
	// File Not Found Error Guard
	if err != nil {
		fmt.Printf("Unable to open file : %v\n", filename);
		os.Exit(1);
	}

	// Buffer Scanner Scans until EOF

	scanner := bufio.NewScanner(envFile);
	numL := 1;
	lastLine := "";

	for scanner.Scan() {
		line := scanner.Text();
		lastLine = line;

		if isComment(line) {
			// Validate Comment
			utils.ValidateComment(line, numL);

		} else if isLine(line) {
			// Key-Value Pair
			utils.ValidatePairline(line, numL);

		}  else if len(line) == 0 {
			// Space (do nothing)
			_ = 0;

		} else {
			// Undefined
			var err *utils.Error = utils.NewError(
				"UndefinedLine",
				"<undefined line could not be parsed>",
				numL,
			);
			utils.Error_Stack.Push(*err);
		}
		
		numL++;
	}

	if len(lastLine) == 0 {
		var err *utils.Error = utils.NewError(
			"EndingBlankLine",
			"<file ends in blank link>",
			numL,
		);
		utils.Error_Stack.Push(*err);
	}

	utils.Error_Stack.Print();
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#");
}

func isLine(line string) bool {
	return !isComment(line) && strings.Contains(line, "=");
}