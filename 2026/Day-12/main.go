package main

import (
	"fmt"
	"iter"
	"os"
	"strings"
)

func ReadInput(path string) iter.Seq[string] {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	return strings.SplitSeq(strings.TrimSuffix(string(data), "\n"), "\n")
}

func usage() {
	msg := `Usage:
	
	\<basecommand\> \<part1 | part2 | hello | help\>

basecommand could be the compiled binary or 'go run main.go'

Please always add a valid a argument to the basecommand

Errors:
	exit code 2 means input error, exit code 1 means code error.
	`

	fmt.Println(msg)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "part1":
	case "part2":
	case "hello":
		fmt.Println("Advent of Code 2026!")
	case "help":
		usage()
	default:
		usage()
		os.Exit(2)
	}
}
