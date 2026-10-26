package main

import (
	"fmt"
	"sg/ghostops/analyzer"
	"sg/ghostops/models/incident"
	"sg/ghostops/parser"
)

func main() {
	//fmt.Println(parser.ParseLine("2026-10-04T19:52:15Z INFO auth User successfully logged in"))
	log, err := parser.ParseLine("2026-10-01T17:50:22Z ERROR payment-service Payment gateway timeout - transaction aborted")
	if err == nil {
		fmt.Println(incident.MatchIncident(*log))
	}

	result, err := analyzer.FileAnalysis("sample.log")

	if err == nil {
		fmt.Println("something went wrong")
	}

	result.ShowResultCL()
}
