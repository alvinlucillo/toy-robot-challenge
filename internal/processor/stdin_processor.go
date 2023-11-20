package processor

import (
	"bufio"
	"fmt"
	"io"
)

type StdinProcessor struct{}

func (StdinProcessor) Process(r io.Reader) error {

	fmt.Println("Welcome to the Toy Robot Challenge Program! 🤖")
	fmt.Println("Enter your commands below:")
	fmt.Println()
	scanner := bufio.NewScanner(r)

	for {
		if scanner.Scan() {
			command := scanner.Text()

			fmt.Println("> You entered ", command)
		}

		if scanner.Err() != nil {
			return scanner.Err()
		}
	}
}
