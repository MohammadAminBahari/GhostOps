package main

import (
	"fmt"
	"log"
	"os"
	"sg/ghostops/analyzer"

)

func main() {
	//args := os.Args[1:]

	//fmt.Println(parser.ParseLine("2026-10-04T19:52:15Z INFO auth User successfully logged in"))
	// log, err := parser.ParseLine("2026-10-01T17:50:22Z ERROR payment-service Payment gateway timeout - transaction aborted")
	// if err == nil {
	// 	fmt.Println(incident.MatchIncident(*log))
	// }

	// result, err := analyzer.FileAnalysis("sample.log")

	// if err == nil {
	// 	fmt.Println(err.Error())
	// }

	// result.ShowResultCL()


	// i used ai help for this section (first i tried to implement this section with flag package but because of the command 
	// it wasn't the best choice so i chose this way)

	args := os.Args[1:]

	var inputFile, outputFile string

	if len(args) > 0 {
		inputFile = args[0]
	}

	for i := 1; i < len(args); i++ {
		if args[i] == "-o" {
			if i+1 >= len(args) {
				log.Fatal("Missing filename after -o")
			}
			outputFile = args[i+1]
			break
		}
	}

	fmt.Println(inputFile)

	result, err := analyzer.FileAnalysis(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	result.ShowResultCL()

	if outputFile != "" {
		if err := result.WriteJson(outputFile); err != nil {
			log.Fatal(err)
		}
		fmt.Println("JSON report saved to", outputFile)
	}
}
