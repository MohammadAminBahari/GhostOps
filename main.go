package main

import (
	"fmt"
	"sg/ghostops/analyzer"
	//"sg/ghostops/parser"
)

func main() {
	//fmt.Println(parser.ParseLine("2026-10-04T19:52:15Z INFO auth User successfully logged in"))


	result, err := analyzer.FileAnalysis("sample.log")

	if err == nil{
		fmt.Println("something went wrong")
	}

	result.ShowResultCL()
}
