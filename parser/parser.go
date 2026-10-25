package parser

import (
	"sg/ghostops/model"
	"sg/ghostops/utils"
	"strings"
	"time"

)

func ParseLine(line string) (*model.LogEntry, error) {

	if len(line) < 4 {
		return nil, MalformedEntryErr
	}

	sections := strings.Fields(line)

	convTime, err := time.Parse(time.RFC3339, sections[0])

	if err != nil {
		return nil, InvalidTimeErr
	}


	convSev, ok := model.SeverityMap[sections[1]]

	if !ok{
		return nil, InvalidSeverityErr
	}

	convSer, ok := utils.ExtractServiceName(sections[2])

	if !ok{
		return nil, InvalidServiceErr
	}






	entry := model.LogEntry{
		TimeStamp: convTime,
		Severity: convSev,
		Service: convSer,
		Message: strings.Join(sections[3:]," "),
	
	}



	return &entry, nil

}
