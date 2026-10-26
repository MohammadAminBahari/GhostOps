package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"sg/ghostops/models/logentry"
	"sg/ghostops/parser"
	"sg/ghostops/models/report"
)

func FileAnalysis(filePath string) (*report.Report, error) {
	report := report.New()
	file, err := os.Open("sample.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		entry, err := parser.ParseLine(line)
		if err != nil {
			report.Stat.MalformedEntries++
		} else {
			report.Stat.Service[entry.Service]++
			report.Stat.Severity[logentry.SeverityToString[entry.Severity]]++
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return report, nil
}
