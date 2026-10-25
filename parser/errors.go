package parser

import "fmt"

var MalformedEntryErr = fmt.Errorf("log entry is not valid")
var InvalidTimeErr = fmt.Errorf("time format is not valid")
var InvalidSeverityErr = fmt.Errorf("severity is not valid")
var InvalidServiceErr = fmt.Errorf("invalid service format")
