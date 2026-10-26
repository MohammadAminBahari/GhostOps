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
	}
}

func (r Report)WriteJson(fileName string){

}

func (r Report)ShowResultCL(){
	r.Stat.ShowCL()
	for _, v := range r.Incident{
		v.ShowCL()
	}
}
