package robot

type Robot interface {
	Place(x, y, direction int) error
	// Move() error
	// Left() error
	// Right() error
	// Report() error
}

type ToyRobot struct {
	state RobotState
}

type RobotState struct {
	x         int
	y         int
	direction int
	isPlaced  bool
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
