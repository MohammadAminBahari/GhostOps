package incident

import (
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
	
}


func MatchIncident(log logentry.LogEntry) *Incident{

	return nil
}
