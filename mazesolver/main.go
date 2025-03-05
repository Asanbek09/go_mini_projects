package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)
}

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}