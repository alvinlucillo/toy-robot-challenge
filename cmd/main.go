package main

import (
	"fmt"
	"os"

	processor "alvinlucillo/toy-robot-challenge/internal/processor"
)

const (
	_ int = iota
	NORTH
	EAST
	WEST
	SOUTH
)

type Direction struct {
	target string
	value  int
}

func main() {
	processor, err := processor.NewProcessor(os.Stdin, processor.SourceTypeStdin)
	if err != nil {
		fmt.Println(err)

		return
	}
	processor.Execute()
}
