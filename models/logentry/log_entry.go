package logentry

import "time"




type SeverityEnum int

const (
	INFO = iota + 1
	ERROR
	WARN
	FATAL
)

type LogEntry struct {
	TimeStamp time.Time    `json:"timestamp"`
	Severity  SeverityEnum `json:"severity"`
	Service   string       `json:"service"`
	Message   string       `json:"message"`
}

var SeverityMap = map[string]SeverityEnum{
	"INFO":          INFO,
	"WARN":          WARN,
	"ERROR":         ERROR,
	"FATAL":         FATAL,
}

var SeverityToString = map[SeverityEnum]string{
	INFO:          "INFO",
	WARN:          "WARN",
	ERROR:         "ERROR",
	FATAL:         "FATAL",
}
