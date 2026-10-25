package main

import (
	"fmt"
	"sg/ghostops/parser"
)

func main() {
	fmt.Println(parser.ParseLine("2026-10-04T19:52:15Z INFO auth User successfully logged in"))
}
