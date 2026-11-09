package main

import (
	"fmt"
	"os"

	"day-9/reader"

	part1 "day-9/part-1"
	part2 "day-9/part-2"
)

func usage() {
	msg := `Usage:
	
	<basecommand> <part1 | part2 | hello | help>

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
		part1Solution := part1.Solve(reader.ReadInput("input/puzzle.txt"))
		fmt.Printf("Part 1 solution: %v\n", part1Solution)
	case "part2":
		part2Solution := part2.Solve(reader.ReadInput("input/puzzle.txt"))
		fmt.Printf("Part 2 solution: %v\n", part2Solution)
	case "hello":
		fmt.Println("Advent of Code 2026!")
	case "help":
		usage()
	default:
		usage()
		os.Exit(2)
	}
}
