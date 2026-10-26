package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"sg/ghostops/models/incident"
	"sg/ghostops/models/logentry"
	"sg/ghostops/models/report"
	"sg/ghostops/parser"
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
		report.Stat.TotalEntries++
		line := scanner.Text()
		//fmt.Println(line)

		entry, err := parser.ParseLine(line)
		if err != nil {
			report.Stat.MalformedEntries++
			continue
		} else {
			report.Stat.Service[entry.Service]++
			report.Stat.Severity[logentry.SeverityToString[entry.Severity]]++
		}

		if entry.Severity != logentry.ERROR && entry.Severity != logentry.FATAL {
			continue
		}

		inc := incident.MatchIncident(*entry)

		if inc == nil {
			continue
		}

		oldInc, ok := report.GetIncidentByTypeService(inc.Type, inc.AffectedService)
		if !ok {
			report.Incident = append(report.Incident, *inc)
		} else {
			oldInc.LastSeen = inc.LastSeen
			oldInc.Frequency++
			exist := report.DoesSeverityExist(logentry.SeverityToString[inc.Severities[0]])
			if !exist {
				oldInc.Severities = append(oldInc.Severities, inc.Severities...)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return report, nil
}
