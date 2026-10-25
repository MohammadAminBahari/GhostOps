package analyzer

import (
	"bufio"
	"fmt"
	"os"
)

func FileAnalysis(filePath string) {
	file, err := os.Open("sample.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// Process the current line here
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}