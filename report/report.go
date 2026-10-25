package report

import "sg/ghostops/model"

type Report struct {
	stat model.Stat
	incident []model.Incident
}


func (r Report)WriteJson(fileName string){

}

func (r Report)ShowResultCL(){

}