package report

import (
	"sg/ghostops/models/incident"
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

func (r Report)WriteJson(fileName string){

}

func (r Report)ShowResultCL(){
	r.Stat.ShowCL()
	for _, inc := range r.Incident{
		inc.ShowCL()
	}
}

func (r *Report)GetIncidentByTypeService(name string, service string) (*incident.Incident, bool){
	// TODO

	return nil, true
}

func (r *Report)DoesSeverityExist(severity string) bool{
	// TODO
	return true
}
