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
	Init(source io.Reader, robot robot.Robot, logger Logger)
	Process() error
}

type Processor struct {
	SrcProcessor SourceProcessor
}

func NewProcessor(source io.Reader, sourceType string) (*Processor, error) {
	var sourceProcessor SourceProcessor
	if sourceType == SourceTypeStdin {
		sourceProcessor = &StdinProcessor{}
	} else {
		return nil, fmt.Errorf("unsupported source type: %s", sourceType)
	}

	sourceProcessor.Init(source, &robot.ToyRobot{}, &StdLogger{})

	return &Processor{
		SrcProcessor: sourceProcessor,
	}, nil
}

func (p *Processor) Execute() {
	p.SrcProcessor.Process()
}
