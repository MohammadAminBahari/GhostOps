package stat

type Stat struct {
	Service          map[string]int `json:"service_stats"`
	Severity         map[string]int `json:"severity_stats"`
	MalformedEntries int            `json:"malformed_entries"`
	TotalEntries     int            `json:"total_entries"`
}

func (stat Stat)ShowCL(){
	
}