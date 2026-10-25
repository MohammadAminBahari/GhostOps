package utils

import "strings"

// for helper functions

func ExtractServiceName(service string) (string, bool) {
	return strings.TrimRight(service, "-service"), strings.HasSuffix(service, "-service") &&
		len(service) > len("-service")
}


