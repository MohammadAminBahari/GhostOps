package model

import "time"

type Stat struct {
	Service          map[string]int `json:"service_stats"`
	Severity         map[string]int `json:"severity_stats"`
	MalformedEntries int            `json:"malformed_entries"`
	TotalEntries     int            `json:"total_entries"`
}

type Incident struct {
	Type            string    `json:"type"`
	Severities      []string  `json:"severity"`
	Frequency       int       `json:"frequency"`
	AffectedService string    `json:"affected_service"`
	FirstSeen       time.Time `json:"first_seen"`
	LastSeen        time.Time `json:"last_seen"`
}

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
	"INALID_LEVEL": INVALID_LEVEL,
	"INFO":         INFO,
	"WARN":         WARN,
	"ERROR":        ERROR,
	"FATAL":        FATAL,
}
