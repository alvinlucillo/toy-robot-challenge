package processor

import (
	"alvinlucillo/toy-robot-challenge/internal/robot"
	"fmt"
	"strings"
	"testing"
)

type MockLogger struct {
	logs []string
}

func (m *MockLogger) Println(args ...interface{}) {
	var log string
	for _, arg := range args {
		log += fmt.Sprint(arg)
	}

	m.logs = append(m.logs, log)
}

func TestProcessData(t *testing.T) {
	stdin := strings.NewReader("PLACE")

	processor := &StdinProcessor{}
	robot := &robot.ToyRobot{}
	logger := &MockLogger{}

	processor.Init(stdin, robot, logger)
	processor.Process()
	fmt.Println(robot.GetState())

	// t.Logf("%s", logger.logs)

	// t.Fail()

}
