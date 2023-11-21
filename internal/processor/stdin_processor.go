package processor

import (
	"alvinlucillo/toy-robot-challenge/internal/robot"
	"bufio"
	"io"
	"strconv"
	"strings"
)

type StdinProcessor struct {
	source io.Reader
	robot  robot.Robot
	logger Logger
}

func (p *StdinProcessor) Init(src io.Reader, robot robot.Robot, logger Logger) {
	p.source = src
	p.robot = robot
	p.logger = logger
}

func (p *StdinProcessor) Process() error {

	p.robot.Init()

	p.logger.Println("Welcome to the Toy Robot Challenge Program! 🤖")
	p.logger.Println("Need help? Enter HELP")
	p.logger.Println("\nEnter your commands below:")

	helpText := `Usage: COMMAND [ARGS]
    Commands:
     PLACE x,y,z  Places the robot on the table in position (x,y) facing z direction 
            - Args:   x - the X coordinate (valid values: 0-3)
                      y - the Y coordinate (valid values: 0-3)
                      z - the direction the robot is facing (valid values: NORTH,EAST,SOUTH,WEST)
            - Example: PLACE 1,2,north
     LEFT         Rotates the robot 90 degrees to the left
     RIGHT        Rotates the robot 90 degrees to the right
     MOVE         Moves the robot one unit forward
     REPORT       Prints the location (X,Y) and direction it's facing`

	scanner := bufio.NewScanner(p.source)

	commandsMap := map[string]bool{
		robot.CommandPlace:  true,
		robot.CommandLeft:   true,
		robot.CommandRight:  true,
		robot.CommandMove:   true,
		robot.CommandReport: true,
		robot.CommandHelp:   true,
	}

	directionsMap := map[string]int{
		"NORTH": robot.DirectionNorth,
		"EAST":  robot.DirectionEast,
		"WEST":  robot.DirectionWest,
		"SOUTH": robot.DirectionSouth,
	}

	for {
		if scanner.Scan() {
			commandParts := strings.Split(strings.TrimSpace(scanner.Text()), " ")

			if found := commandsMap[commandParts[0]]; !found {
				p.logger.Println("> Invalid command. Enter HELP for usage.")
				continue
			}

			if commandParts[0] == robot.CommandPlace {
				p.robot.Place(1, 2, robot.DirectionNorth)
			} else {
				switch commandParts[0] {
				case robot.CommandHelp, strings.ToLower(robot.CommandHelp):
					p.logger.Println(helpText)
				case robot.CommandLeft, robot.CommandRight, robot.CommandMove, robot.CommandReport:
					if !p.robot.IsPlaced() {
						p.logger.Println("> Oops. Robot not yet placed. Enter a PLACE command first.")
						continue
					}
					fallthrough
				case robot.CommandPlace:
					if len(commandParts) != 2 {
						p.logger.Println("> Invalid use of PLACE. Enter HELP for usage.")
						continue
					}

					placeParts := strings.Split(commandParts[1], ",")
					if len(placeParts) != 3 {
						p.logger.Println("> Invalid use of PLACE. Enter HELP for usage.")
						continue
					}

					xValue, err := strconv.Atoi(placeParts[0])
					if err != nil {
						p.logger.Println("> Invalid use of PLACE. Enter HELP for usage.")
						continue
					}

					yValue, err := strconv.Atoi(placeParts[1])
					if err != nil {
						p.logger.Println("> Invalid use of PLACE. Enter HELP for usage.")
						continue
					}

					direction := directionsMap[placeParts[2]]
					if direction == 0 {
						p.logger.Println("> Invalid use of PLACE. Enter HELP for usage.")
						continue
					}

					p.robot.Place(xValue, yValue, direction)
				}
			}
		} else {
			if scanner.Err() != nil {
				return scanner.Err()
			}

			break
		}
	}

	return nil
}
