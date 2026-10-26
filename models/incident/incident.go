package incident

import (
	"fmt"
	"sg/ghostops/models/logentry"
	//"sg/ghostops/utils"
	"time"
)

type Incident struct {
	Type            string    `json:"type"`
	Severities      []string  `json:"severity"`
	Frequency       int       `json:"frequency"`
	AffectedService string    `json:"affected_service"`
	FirstSeen       time.Time `json:"first_seen"`
	LastSeen        time.Time `json:"last_seen"`
}

func (inc Incident)ShowCL(){
	fmt.Println(inc.Type)
	for _, sev := range inc.Severities{
		fmt.Printf("%s, ", sev)
	}

	fmt.Printf("Number of Accurance: %d\n on %s Service\n", inc.Frequency, inc.AffectedService)
	fmt.Printf("First Accurance at: %s\n Last Accurance at: %s", inc.FirstSeen.String(), inc.LastSeen.String())


}


func MatchIncident(log logentry.LogEntry) *Incident{

	return nil
}
