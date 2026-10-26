package utils

import (
	"strings"
	"time"
)

// for helper functions

func ExtractServiceName(service string) (string, bool) {
	return strings.TrimSuffix(service, "-service"), strings.HasSuffix(service, "-service") &&
		len(service) > len("-service")
}

func GetEarliestTime(t1 time.Time, t2 time.Time) time.Time{
	if t1.Before(t2) {return t1}
	return t2
}

func GetLatestTime(t1 time.Time, t2 time.Time) time.Time{
	if t1.After(t2) {return t1}
	return t2
}

