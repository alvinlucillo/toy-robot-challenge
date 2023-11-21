package processor

import "fmt"

type Logger interface {
	// Log(args ...interface{})
	Println(args ...interface{})
}

type StdLogger struct{}

func (s *StdLogger) Println(args ...interface{}) {
	fmt.Println(args...)
}
