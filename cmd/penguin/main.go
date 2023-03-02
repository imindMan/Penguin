/*
 * This is a project to develop a kid-friendly programming language called Penguin.
 * All necessary information can be found in the GitHub repository or read along for more details.
 */

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if an input file was provided
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No input file provided\n")
	}

	// Open the input file
	inputFile := args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// TODO: implement functionality
}
