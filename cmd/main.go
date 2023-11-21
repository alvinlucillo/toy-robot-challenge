package main

import (
	"fmt"
	"os"

	processor "alvinlucillo/toy-robot-challenge/internal/processor"
)

func main() {
	processor, err := processor.NewProcessor(os.Stdin, processor.SourceTypeStdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	processor.Execute()
}
