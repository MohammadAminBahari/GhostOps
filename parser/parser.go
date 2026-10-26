package parser

import (
	"sg/ghostops/models/logentry"
	"sg/ghostops/utils"
	"strings"
	"time"

)

func ParseLine(line string) (*logentry.LogEntry, error) {

	

	sections := strings.Fields(line)

	if len(sections) < 4 {
		return nil, MalformedEntryErr
	}

	convTime, err := time.Parse(time.RFC3339, sections[0])

	if err != nil {
		return nil, InvalidTimeErr
	}


	convSev, ok := logentry.SeverityMap[sections[1]]

	if !ok{
		return nil, InvalidSeverityErr
	}

	convSer, ok := utils.ExtractServiceName(sections[2])

	if !ok{
		return nil, InvalidServiceErr
	}






	entry := logentry.LogEntry{
		TimeStamp: convTime,
		Severity: convSev,
		Service: convSer,
		Message: strings.Join(sections[3:]," "),
	
	}



	return &entry, nil

}
