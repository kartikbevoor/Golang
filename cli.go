package main

import (
	"flag"
	"fmt"
	"os"
)

// A CLI (Command Line Interface) program is an application that runs in the terminal
// and interacts with users through text commands and arguments.

// CLI Program in Go: Runs from terminal, Accepts arguments, Parses flags, Executes logic.
func cli() {

	// Reading command line arguments
	fmt.Println(os.Args) // os.Args[0] → program name, os.Args[1] → first argument, os.Args[2] → second argument

	simpleGreeter()

}

func simpleGreeter() {
	if len(os.Args) < 2 {
		fmt.Println("Enter your name")
	} else {
		name := os.Args[1]
		fmt.Println("Hello", name)
	}
}

// Flags are command-line parameters passed when running a program.
func usingFlags() {
	names := flag.String("Name", "guest", "your name") // arguments: flag name, degault value, description
	age := flag.Int("Age", 18, "vayasu")

	flag.Parse()

	fmt.Println("Name:", *names)
	fmt.Println("Age:", *age)
} // go run main.go -name=Kartik -age=21

// Creating subcommands
func createSubcommands() {
	if len(os.Args) < 2 {
		fmt.Println("Expected arguments")
		return
	}

	switch os.Args[1] {
	case "create":
		fmt.Println("Create")
	case "delete":
		fmt.Println("Delete")
	}
}
