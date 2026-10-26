package logentry

import "time"




type SeverityEnum int

const (
	INVALID_LEVEL = iota
	INFO
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
	"INVALID_LEVEL": INVALID_LEVEL,
	"INFO":          INFO,
	"WARN":          WARN,
	"ERROR":         ERROR,
	"FATAL":         FATAL,
}

var SeverityToString = map[SeverityEnum]string{
	INVALID_LEVEL: "INVALID_LEVEL",
	INFO:          "INFO",
	WARN:          "WARN",
	ERROR:         "ERROR",
	FATAL:         "FATAL",
}
