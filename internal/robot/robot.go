package robot

import "fmt"

type ToyRobot struct {
	state                RobotState
	mapDirectionsByTitle map[string]int
	mapDirectionsByValue map[int]string
}

func (t *ToyRobot) Place(x, y int, direction string) error {
	d := t.mapDirectionsByTitle[direction]

	if (x >= 0 && x <= 3) && (y >= 0 && y <= 3) && d != 0 {
		t.state.x = x
		t.state.y = y
		t.state.direction = d
		t.state.isPlaced = true

		return nil
	} else {
		return fmt.Errorf("invalid robot coordinates or direction")
	}
}

func (t *ToyRobot) Report() string {
	return fmt.Sprintf("Output: %v,%v,%v\n", t.state.x, t.state.y, t.mapDirectionsByValue[t.state.direction])
}

func (t *ToyRobot) Left() {
	t.changeDirection(-1)
}

func (t *ToyRobot) Right() {
	t.changeDirection(1)
}

func (t *ToyRobot) changeDirection(value int) {
	d := (t.state.direction + value) % 4

	if d == 0 {
		d = 4
	}

	t.state.direction = d
}

func (t *ToyRobot) IsPlaced() bool {
	return t.state.isPlaced
}

func (t *ToyRobot) GetState() RobotState {
	return t.state
}

func (t *ToyRobot) Move() error {
	// moving to north means adding one unit to the y coordinate
	// moving to south means subtracting one from the y coordinate
	// etc.
	directionMoveMap := map[int]struct {
		coordinate string
		value      int
	}{
		DirectionNorth: {"y", 1},
		DirectionEast:  {"x", 1},
		DirectionSouth: {"y", -1},
		DirectionWest:  {"x", -1},
	}

	// temporary variable (state) contains modified state
	direction := directionMoveMap[t.state.direction]
	state := t.state
	if direction.coordinate == "x" {
		state.x += direction.value
	} else {
		state.y += direction.value
	}

	// only apply the temporary state if the robot doesn't fall off the table
	if state.x >= 0 && state.x <= 3 && state.y >= 0 && state.y <= 3 {
		t.state = state
		return nil
	} else {
		return fmt.Errorf("the robot falls off the table")
	}
}

func (t *ToyRobot) Init() {
	t.state = RobotState{
		x:         -1,
		y:         -1,
		direction: -1,
	}

	t.mapDirectionsByTitle = map[string]int{
		DirectionNorthTitle: DirectionNorth,
		DirectionEastTitle:  DirectionEast,
		DirectionSouthTitle: DirectionSouth,
		DirectionWestTitle:  DirectionWest,
	}

	t.mapDirectionsByValue = map[int]string{
		DirectionNorth: DirectionNorthTitle,
		DirectionEast:  DirectionEastTitle,
		DirectionSouth: DirectionSouthTitle,
		DirectionWest:  DirectionWestTitle,
	}
}
