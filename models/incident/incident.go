package incident

import (
	"fmt"
	"sg/ghostops/models/logentry"
	"sg/ghostops/utils"
	"strings"
	"time"
)

type Incident struct {
	Type            string                  `json:"type"`
	Severities      []logentry.SeverityEnum `json:"severity"`
	Frequency       int                     `json:"frequency"`
	AffectedService string                  `json:"affected_service"`
	FirstSeen       time.Time               `json:"first_seen"`
	LastSeen        time.Time               `json:"last_seen"`
}

func (inc Incident) ShowCL() {
	fmt.Println(inc.Type)
	for _, sev := range inc.Severities {
		fmt.Printf("%s, ", logentry.SeverityToString[sev])
	}

	fmt.Printf("Number of Accurance: %d\non %s Service\n", inc.Frequency, inc.AffectedService)
	fmt.Printf("First Accurance at: %s\nLast Accurance at: %s\n", inc.FirstSeen.String(), inc.LastSeen.String())

}

func MatchIncident(log logentry.LogEntry) *Incident {
	message := strings.ToLower(log.Message)

	for name, keywords := range utils.IncidentKeywords {
		for _, keyword := range keywords {
			if strings.Contains(message, keyword) {
				return &Incident{
					Type:            name,
					Severities:      []logentry.SeverityEnum{log.Severity},
					AffectedService: log.Service,
					Frequency:       1,
					FirstSeen:       log.TimeStamp,
					LastSeen:        log.TimeStamp,
				}
			}
		}
	}

	return nil
}
