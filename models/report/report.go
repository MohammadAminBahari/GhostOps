package report

import (
	"encoding/json"
	"fmt"
	"os"
	"sg/ghostops/models/incident"
	"sg/ghostops/models/logentry"
	"sg/ghostops/models/stat"
)

type Report struct {
	Stat stat.Stat
	Incident []incident.Incident
}


func New() *Report{
	return &Report{
		Stat: stat.Stat{
			Service: make(map[string]int),
			Severity: make(map[string]int),
		},
		Incident: make([]incident.Incident, 0),
	}
}

func (r Report)WriteJson(fileName string) error{
	data, err := json.MarshalIndent(r, "", "    ")
    if err != nil {
        return err
    }

    return os.WriteFile(fileName, data, 0644)
}

func (r Report)ShowResultCL(){
	fmt.Println("Satistics: ")
	r.Stat.ShowCL()
	fmt.Println("\nIncidents: ")
	for _, inc := range r.Incident{
		inc.ShowCL()
		fmt.Println("-----------------------------------------")
	}

	
}

func (r *Report)GetIncidentByTypeService(name string, service string) (*incident.Incident, bool){

	for  i := range r.Incident{
		if r.Incident[i].AffectedService == service && r.Incident[i].Type == name{
			return &r.Incident[i], true
		}
	}
	return nil, false
}

func (r *Report)DoesSeverityExist(severity string) bool{

	for _, inc := range r.Incident{
		for _, sev := range inc.Severities{
			if logentry.SeverityToString[sev] == severity{
				return true;
			}
		}
	}
	return false
}
