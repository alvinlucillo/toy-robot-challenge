package robot

const (
	_ int = iota
	DirectionNorth
	DirectionEast
	DirectionWest
	DirectionSouth
)

const (
	CommandMove   = "MOVE"
	CommandPlace  = "PLACE"
	CommandLeft   = "LEFT"
	CommandRight  = "RIGHT"
	CommandReport = "REPORT"
	CommandHelp   = "HELP"
)

type RobotState struct {
	x         int
	y         int
	direction int
	isPlaced  bool
}

type Robot interface {
	Init()
	Place(x, y, direction int) error
	// Move() error
	// Left() error
	// Right() error
	Report()
	IsPlaced() bool
	GetState() RobotState
}
