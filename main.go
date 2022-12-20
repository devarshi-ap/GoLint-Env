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
			fmt.Printf("Comment --> %v\n\n", line);
			utils.ValidateComment(line, numL);

		} else if isLine(line) {
			// Validate Key-Value Pair
			fmt.Printf("K-V Pair --> %v\n\n", line);
			utils.ValidatePairline(line, numL);

		}  else if len(line) == 0 {
			// Space (do nothing)
			fmt.Printf("--\n\n")

		} else {
			// Process Undefined
			fmt.Printf("Undefined\n\n");
		}
		
		numL++;
	}

	if len(lastLine) == 0 {
		var err *utils.Error = utils.NewError("", "", 1);
		utils.Error_Stack.Push(*err);
	}

	fmt.Printf("---------------------------------\n\n")
	utils.Error_Stack.Print();
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#");
}

func isLine(line string) bool {
	return !isComment(line) && strings.Contains(line, "=")
}