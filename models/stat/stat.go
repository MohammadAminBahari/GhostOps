package stat

import "fmt"

type Stat struct {
	Service          map[string]int `json:"service_stats"`
	Severity         map[string]int `json:"severity_stats"`
	MalformedEntries int            `json:"malformed_entries"`
	TotalEntries     int            `json:"total_entries"`
}

func (stat Stat) ShowCL() {
	fmt.Printf("Totla Entries: %d \t Malformed Entries: %d \nService Satistics: \n", stat.TotalEntries, stat.MalformedEntries)
	for key, value := range stat.Service {
		fmt.Printf("%s: %d \t", key, value)
	}
	fmt.Printf("\nSeverity Level Satistics: \n")
	for key, value := range stat.Severity {
		fmt.Printf("%s: %d \t", key, value)
	}
}
