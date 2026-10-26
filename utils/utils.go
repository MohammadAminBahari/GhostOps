package utils

import (
	"strings"
	"time"
)

// for helper functions

func ExtractServiceName(service string) (string, bool) {
	return strings.TrimRight(service, "-service"), strings.HasSuffix(service, "-service") &&
		len(service) > len("-service")
}

func CompareTimes(t1 time.Time, t2 time.Time){
	
}

