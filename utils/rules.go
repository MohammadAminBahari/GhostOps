package utils


// i used ai to extract these keywords from the sample.log file 

var IncidentKeywords = map[string][]string{
	"database_connection_failure": {
		"database connection",
		"connection pool exhausted",
	},
	"payment_gateway_timeout": {
		"payment gateway timeout",
		"transaction aborted",
	},
	"authentication_failure": {
		"authentication failed",
		"login failed",
		"unauthorized",
	},
	"network_timeout": {
		"connection timeout",
		"request timeout",
		"i/o timeout",
	},
	"service_crash": {
		"crashed",
		"panic",
		"fatal error",
	},
	"resource_exhaustion": {
		"out of memory",
		"resource exhausted",
		"too many open files",
	},
}