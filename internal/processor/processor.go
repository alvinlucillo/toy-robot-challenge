package processor

import (
	"fmt"
	"io"

	robot "alvinlucillo/toy-robot-challenge/internal/robot"
)

const (
	SourceTypeStdin = "stdin"
	SourceTypeFile  = "file"
)

type SourceProcessor interface {
	Process(reader io.Reader) error
}

type Processor struct {
	SrcProcessor SourceProcessor
	Source       io.Reader
	Robot        robot.Robot
}

func NewProcessor(source io.Reader, sourceType string) (*Processor, error) {
	var sourceProcessor SourceProcessor
	if sourceType == SourceTypeStdin {
		sourceProcessor = &StdinProcessor{}
	} else {
		return nil, fmt.Errorf("unsupported source type: %s", sourceType)
	}

	return &Processor{
		SrcProcessor: sourceProcessor,
		Source:       source,
		Robot:        &robot.ToyRobot{},
	}, nil
}

func (p *Processor) Execute() {
	p.SrcProcessor.Process(p.Source)
}
