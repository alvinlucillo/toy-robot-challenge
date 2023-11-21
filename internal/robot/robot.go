package robot

import "fmt"

type ToyRobot struct {
	state RobotState
}

func (t *ToyRobot) Place(x, y, direction int) error {
	if (x >= 0 && x <= 4) && (y >= 0 && y <= 4) {
		t.state.x = x
		t.state.y = y
		t.state.direction = direction
		t.state.isPlaced = true
	}

	return nil
}

func (t *ToyRobot) Report() {
	fmt.Printf("Output: %v,%v,%v\n", t.state.x, t.state.y, t.state.direction)
}

func (t *ToyRobot) IsPlaced() bool {
	return t.state.isPlaced
}

func (t *ToyRobot) GetState() RobotState {
	return t.state
}

func (t *ToyRobot) Init() {
	t.state = RobotState{
		x:         -1,
		y:         -1,
		direction: -1,
	}
}
